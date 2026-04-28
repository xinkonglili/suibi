package _chan

import (
	"context"
	"fmt"
	"testing"
	"time"
)

/*tick:0,2026-04-28 10:37:03.8686195 +0800 CST m=+1.016512201
tick:0,2026-04-28 10:37:04.868375 +0800 CST m=+2.016267701
tick:0,2026-04-28 10:37:05.868488 +0800 CST m=+3.016380701
tick:0,2026-04-28 10:37:06.8683802 +0800 CST m=+4.016272901
tick:0,2026-04-28 10:37:07.8684256 +0800 CST m=+5.016318301
tick:0,2026-04-28 10:37:08.8683646 +0800 CST m=+6.016257301*/
//1、一直循环，没有退出的条件
//for range tick.C {
//	count := 0
//	fmt.Println(fmt.Sprintf("tick:%d,%v", count, time.Now()))
//	count++
//}
//

// 加上退出的条件
//
//	for {
//		select {
//		case <-tick.C:
//			count := 0
//			fmt.Println(fmt.Sprintf("tick:%d,%v", count, time.Now()))
//			count++
//			if count > 5 {
//				return
//			}
//		default:
//			fmt.Println("default") //死循环，容易cpu 100%
//		}
func TestChanTime(t *testing.T) {

	tick := time.NewTicker(1 * time.Second)
	defer tick.Stop()

	//}
	//加上退出的条件
	var count int
	for {
		select {
		case <-tick.C:
			fmt.Println(fmt.Sprintf("tick:%d,%v", count, time.Now()))
			count++
			if count > 3 {
				return
			}
		}
	}
}

func TestChanTimeWithContext(t *testing.T) {
	// 设置一个 4 秒的超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	tick := time.NewTicker(1 * time.Second)
	defer tick.Stop()

	var count int
	for {
		select {
		case <-ctx.Done(): // 当超时或被取消时触发
			fmt.Println("任务结束:", ctx.Err())
			return
		case <-tick.C:
			count++
			fmt.Printf("tick:%d,%v\n", count, time.Now())
		}
	}
}

// 如果你需要在代码的其他地方手动触发退出，可以使用一个空的 struct 通道。
func TestChanTimeWithSignal(t *testing.T) {
	tick := time.NewTicker(1 * time.Second)
	defer tick.Stop()

	done := make(chan struct{}) // 退出信号通道

	// 模拟在另一个 goroutine 中触发退出
	go func() {
		time.Sleep(4 * time.Second)
		close(done) // 关闭通道作为广播信号
	}()

	var count int
	for {
		select {
		case <-done: // 监听到退出信号
			fmt.Println("收到退出信号")
			return
		case <-tick.C:
			count++
			fmt.Printf("tick:%d,%v\n", count, time.Now())
		}
	}
}

// 使用 time.After (适用于一次性超时)
// 如果你不需要周期性的 ticker，或者只是想给整个循环加一个总超时，可以直接在 select 中使用 time.After。
func TestChanTimeWithAfter(t *testing.T) {
	tick := time.NewTicker(1 * time.Second)
	defer tick.Stop()

	timeout := time.After(4 * time.Second) // 4秒后发送一个时间信号

	var count int
	for {
		select {
		case <-timeout: // 超时触发
			fmt.Println("超时退出")
			return
		case <-tick.C:
			count++
			fmt.Printf("tick:%d,%v\n", count, time.Now())
		}
	}
}
