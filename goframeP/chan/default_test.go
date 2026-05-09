package _chan

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestDefault(t *testing.T) {
	timeContext, cancel := context.WithTimeout(context.Background(), 1*time.Second)

	done := make(chan struct{}) // 创建一个无缓冲的channel，阻塞主程序的结束
	defer cancel()
	go func() {
		for {
			select {
			case <-timeContext.Done():
				fmt.Println("timeContext end")
				close(done)
				return
			default:
				time.Sleep(500 * time.Millisecond)
				fmt.Println("-----no result")
			}
		}
	}()

	<-done

}
