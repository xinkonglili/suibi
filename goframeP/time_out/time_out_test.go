package time_out

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"time"
)

/*
推荐文章
1. Go 中如何正确重试请求（最推荐）
涵盖重试策略、熔断降级、对冲请求，结合 hystrix-go 实践
👉 https://www.luozhiyun.com/archives/677
2. 记 Go 中一次 HTTP 超时引发的事故
真实踩坑案例，讲了错误的重试写法和正确实现，很有参考价值
👉 https://www.cnblogs.com/ricklz/p/14840205.html
3. Go 中 HTTP 超时问题的排查（腾讯云）
讲了高并发下 HTTP1.1 创建大量连接导致超时的问题，以及调大 MaxIdleConnsPerHost 参数缓解的方案
👉 https://cloud.tencent.com/developer/article/1529840 Tencent Cloud
4. 有趣的 Go HttpClient 超时机制（源码分析）
深入分析 Go 底层超时实现原理，适合想深入理解的
👉 https://www.cnblogs.com/zhuochongdashi/p/16893627.html*/
// ========================================================
// Demo 1: 模拟 Go HttpClient 的整体超时机制
// ========================================================
func demoHttpClientTimeout() {
	fmt.Println("===== Demo 1: HttpClient 整体超时机制 =====")

	// 1. 创建一个模拟的慢速 HTTP 服务器，处理请求需要 3 秒
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	}))
	defer server.Close()

	// 2. 创建 HttpClient，设置 1 秒的整体超时 (包含连接、重定向、读取)
	// 相比于分别设置连接超时和读取超时，这种方式对使用者更友好
	client := http.Client{
		Timeout: 1 * time.Second,
	}

	resp, err := client.Get(server.URL)
	if err != nil {
		fmt.Printf("请求失败(预期内): %v\n\n", err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("请求成功: %s\n\n", body)
}

// ========================================================
// Demo 2: 模拟底层原理 (协程 + Channel + Select + Context)
// 演示文章中提到的“循环任务”与“阻塞任务”超时控制套路
// ========================================================
func demoUnderlyingPrinciple() {
	fmt.Println("===== Demo 2: 底层原理模拟 (Context + Goroutine + Select) =====")

	// 设置 1 秒超时的 Context
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel() // 释放资源

	fmt.Println("1. 演示循环任务的超时退出:")
	loopTask(ctx)

	fmt.Println("\n2. 演示阻塞任务的超时退出:")
	blockingTask(ctx)
}

// 循环任务套路：子协程起循环，每次循环 select ctx.Done()
func loopTask(ctx context.Context) {
	ch := make(chan int)

	go func(ctx context.Context) {
		n := 1
		for {
			select {
			case <-ctx.Done():
				fmt.Println("   -> 循环任务收到取消信号，退出协程防止泄漏")
				return
			case ch <- n:
				n++
				time.Sleep(300 * time.Millisecond) // 模拟耗时
			}
		}
	}(ctx)

	// 主协程消费数据
	for val := range ch {
		fmt.Printf("   -> 收到数据: %d\n", val)
	}
}

// 阻塞任务套路：子协程 select 阻塞任务与 ctx.Done()
func blockingTask(ctx context.Context) {
	ch := make(chan string)

	go func(ctx context.Context) {
		// 模拟一个耗时的阻塞任务（比如 HTTP 连接建立）
		result := simulateSlowBlockingOp()
		select {
		case <-ctx.Done():
			fmt.Println("   -> 阻塞任务完成前收到取消信号，退出协程")
			return
		case ch <- result:
			fmt.Println("   -> 阻塞任务正常完成，发送结果")
		}
	}(ctx)

	select {
	case <-ctx.Done():
		fmt.Printf("   -> 主协程等待超时: %v\n", ctx.Err())
	case res := <-ch:
		fmt.Printf("   -> 主协程收到结果: %s\n", res)
	}
}

// 模拟耗时 2 秒的阻塞操作
func simulateSlowBlockingOp() string {
	time.Sleep(2 * time.Second)
	return "slow-op-result"
}

func main() {
	demoHttpClientTimeout()
	time.Sleep(500 * time.Millisecond)
	demoUnderlyingPrinciple()
}
