package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

// ─── 数据结构 ───────────────────────────────────────────────

type Message struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Payload   string    `json:"payload"`
	Status    string    `json:"status"` // queued | consuming | acked | nacked
	CreatedAt time.Time `json:"createdAt"`
	Consumer  string    `json:"consumer,omitempty"`
}

type Consumer struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Icon   string `json:"icon"`
	Status string `json:"status"` // idle | busy
}

type Stats struct {
	Sent     int `json:"sent"`
	Queued   int `json:"queued"`
	Consumed int `json:"consumed"`
	Acked    int `json:"acked"`
	Nacked   int `json:"nacked"`
}

type Event struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

// ─── 消息队列核心 ─────────────────────────────────────────

type MQBroker struct {
	mu        sync.Mutex
	queue     []*Message
	history   []*Message
	consumers []*Consumer
	stats     Stats
	clients   map[chan Event]bool
}

func NewMQBroker() *MQBroker {
	b := &MQBroker{
		clients: make(map[chan Event]bool),
		consumers: []*Consumer{
			{ID: "c1", Name: "短信通知", Icon: "📱", Status: "idle"},
			{ID: "c2", Name: "更新库存", Icon: "📦", Status: "idle"},
			{ID: "c3", Name: "通知司机", Icon: "🚗", Status: "idle"},
		},
	}
	// 后台自动消费协程
	go b.autoConsumeLoop()
	return b
}

func (b *MQBroker) Publish(msgType, payload string) *Message {
	b.mu.Lock()
	defer b.mu.Unlock()

	msg := &Message{
		ID:        fmt.Sprintf("msg_%d", time.Now().UnixMilli()),
		Type:      msgType,
		Payload:   payload,
		Status:    "queued",
		CreatedAt: time.Now(),
	}
	b.queue = append(b.queue, msg)
	b.history = append(b.history, msg)
	b.stats.Sent++
	b.stats.Queued++

	b.broadcast(Event{Type: "message_queued", Payload: msg})
	b.broadcast(Event{Type: "stats", Payload: b.stats})
	b.broadcast(Event{Type: "queue_snapshot", Payload: b.queueSnapshot()})
	return msg
}

func (b *MQBroker) queueSnapshot() []string {
	ids := make([]string, len(b.queue))
	for i, m := range b.queue {
		ids[i] = m.ID
	}
	return ids
}

func (b *MQBroker) autoConsumeLoop() {
	for {
		time.Sleep(time.Duration(800+rand.Intn(600)) * time.Millisecond)
		b.mu.Lock()
		if len(b.queue) == 0 {
			b.mu.Unlock()
			continue
		}
		// 找空闲 consumer
		var free *Consumer
		for _, c := range b.consumers {
			if c.Status == "idle" {
				free = c
				break
			}
		}
		if free == nil {
			b.mu.Unlock()
			continue
		}

		msg := b.queue[0]
		b.queue = b.queue[1:]
		msg.Status = "consuming"
		msg.Consumer = free.Name
		free.Status = "busy"
		b.stats.Queued--

		b.broadcast(Event{Type: "message_consuming", Payload: msg})
		b.broadcast(Event{Type: "consumer_update", Payload: free})
		b.broadcast(Event{Type: "queue_snapshot", Payload: b.queueSnapshot()})
		b.mu.Unlock()

		// 模拟处理耗时
		go func(m *Message, c *Consumer) {
			delay := time.Duration(400+rand.Intn(800)) * time.Millisecond
			time.Sleep(delay)

			b.mu.Lock()
			defer b.mu.Unlock()

			// 5% 概率 NACK
			if rand.Intn(20) == 0 {
				m.Status = "nacked"
				b.stats.Nacked++
				// NACK 重新入队
				b.queue = append([]*Message{m}, b.queue...)
				b.stats.Queued++
				b.broadcast(Event{Type: "message_nacked", Payload: m})
			} else {
				m.Status = "acked"
				b.stats.Consumed++
				b.stats.Acked++
				b.broadcast(Event{Type: "message_acked", Payload: m})
			}
			c.Status = "idle"
			b.broadcast(Event{Type: "consumer_update", Payload: c})
			b.broadcast(Event{Type: "stats", Payload: b.stats})
			b.broadcast(Event{Type: "queue_snapshot", Payload: b.queueSnapshot()})
		}(msg, free)
	}
}

func (b *MQBroker) ClearQueue() {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.queue = nil
	b.stats.Queued = 0
	b.broadcast(Event{Type: "stats", Payload: b.stats})
	b.broadcast(Event{Type: "queue_snapshot", Payload: []string{}})
}

func (b *MQBroker) GetHistory() []*Message {
	b.mu.Lock()
	defer b.mu.Unlock()
	result := make([]*Message, len(b.history))
	copy(result, b.history)
	return result
}

func (b *MQBroker) GetStats() Stats {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.stats
}

func (b *MQBroker) GetConsumers() []*Consumer {
	b.mu.Lock()
	defer b.mu.Unlock()
	result := make([]*Consumer, len(b.consumers))
	copy(result, b.consumers)
	return result
}

// ─── SSE 广播 ─────────────────────────────────────────────

func (b *MQBroker) Subscribe() chan Event {
	ch := make(chan Event, 32)
	b.mu.Lock()
	b.clients[ch] = true
	b.mu.Unlock()
	return ch
}

func (b *MQBroker) Unsubscribe(ch chan Event) {
	b.mu.Lock()
	delete(b.clients, ch)
	b.mu.Unlock()
}

func (b *MQBroker) broadcast(evt Event) {
	for ch := range b.clients {
		select {
		case ch <- evt:
		default:
		}
	}
}

// ─── HTTP 处理器 ──────────────────────────────────────────

var broker = NewMQBroker()

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next(w, r)
	}
}

func handlePublish(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", 405)
		return
	}
	var body struct {
		Type    string `json:"type"`
		Payload string `json:"payload"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "bad request", 400)
		return
	}
	msg := broker.Publish(body.Type, body.Payload)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}

func handleHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(broker.GetHistory())
}

func handleStats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(broker.GetStats())
}

func handleConsumers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(broker.GetConsumers())
}

func handleClear(w http.ResponseWriter, r *http.Request) {
	broker.ClearQueue()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

// SSE 长连接推送
func handleSSE(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "streaming not supported", 500)
		return
	}

	ch := broker.Subscribe()
	defer broker.Unsubscribe(ch)

	// 推送初始状态
	sendSSE(w, flusher, Event{Type: "stats", Payload: broker.GetStats()})
	sendSSE(w, flusher, Event{Type: "consumers", Payload: broker.GetConsumers()})
	sendSSE(w, flusher, Event{Type: "history", Payload: broker.GetHistory()})

	ctx := r.Context()
	for {
		select {
		case <-ctx.Done():
			return
		case evt := <-ch:
			sendSSE(w, flusher, evt)
		}
	}
}

func sendSSE(w http.ResponseWriter, f http.Flusher, evt Event) {
	data, _ := json.Marshal(evt)
	fmt.Fprintf(w, "data: %s\n\n", data)
	f.Flush()
}

func main() {
	http.HandleFunc("/api/publish", corsMiddleware(handlePublish))
	http.HandleFunc("/api/history", corsMiddleware(handleHistory))
	http.HandleFunc("/api/stats", corsMiddleware(handleStats))
	http.HandleFunc("/api/consumers", corsMiddleware(handleConsumers))
	http.HandleFunc("/api/clear", corsMiddleware(handleClear))
	http.HandleFunc("/api/events", handleSSE)

	log.Println("🚀 MQ Backend running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
