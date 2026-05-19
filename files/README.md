# MQ Demo — Go 后端 + 原生 HTML 前端

实现了一个完整的消息队列演示系统，包含：
- Go 后端（无第三方依赖）
- SSE 实时推送（Server-Sent Events）
- 生产者发布、队列缓冲、多消费者消费、ACK/NACK 机制

## 项目结构

```
mq-demo/
├── backend/
│   ├── main.go     # Go 后端，内置 MQ broker + HTTP API + SSE
│   └── go.mod
└── frontend/
    └── index.html  # 前端控制台，纯 HTML/CSS/JS
```

## 快速启动

### 1. 启动后端

```bash
cd backend
go run main.go
# 输出: 🚀 MQ Backend running on :8080
```

### 2. 打开前端

直接用浏览器打开 `frontend/index.html`，或起一个静态服务：

```bash
cd frontend
npx serve .
# 或
python3 -m http.server 3000
```

## API 接口

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /api/publish | 发布消息 `{"type":"order","payload":"..."}` |
| GET  | /api/history | 获取消息历史 |
| GET  | /api/stats   | 获取统计数据 |
| GET  | /api/consumers | 获取消费者状态 |
| POST | /api/clear   | 清空队列 |
| GET  | /api/events  | SSE 实时事件流 |

## 核心机制

- **持久化缓冲**：消息先进队列，消费者下线期间消息不丢失
- **手动 ACK**：消费者处理完才确认，模拟 5% NACK 重新入队
- **并发安全**：sync.Mutex 保护队列操作
- **SSE 推送**：后端状态实时广播到所有前端连接
- **多消费者**：3 个独立消费者并发消费，互不影响

## 替换为真实 RabbitMQ

将 `MQBroker` 替换为 `github.com/rabbitmq/amqp091-go` 即可对接真实 MQ，
API 层和前端无需改动。
