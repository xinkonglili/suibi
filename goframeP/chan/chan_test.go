package _chan

import (
	"fmt"
	"testing"
)

func TestChan(t *testing.T) {
	chanA()
	chanB()
	chanC()
}

//chan：
//1、注意：chan 默认是阻塞的，如果chan没有被消费，那么写入chan的数据就会阻塞，直到chan被消费
//2、注意有缓冲区，如果缓冲区已满，那么写入chan的数据就会阻塞，直到缓冲区有空间
//3、向已关闭的chan写入数据会panic，读数据不会 panic

func chanA() {

	ch := make(chan int)

	fmt.Println("test -11")
	ch <- 1
	fmt.Println("test -1") //不会输出，卡住
	ch <- 2
	fmt.Println("test 0")
	ch <- 3
	ch <- 4
	close(ch)
	for c := range ch {
		fmt.Println(c)
	}
}

func chanB() {

	ch := make(chan int, 10)

	fmt.Println("test -11")
	ch <- 1
	fmt.Println("test -1")
	ch <- 2
	fmt.Println("test 0")
	ch <- 3
	ch <- 4
	close(ch)
	for c := range ch {
		fmt.Println(c)
	}
}

func chanC() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	ch := make(chan int, 10)

	fmt.Println("test -11")
	ch <- 1
	fmt.Println("test -1")
	ch <- 2
	fmt.Println("test 0")
	ch <- 3
	ch <- 4
	close(ch)
	ch <- 5
}
