package goroutine

import (
	"fmt"
	"time"
)

/*
1. 关闭一个为初始化的 channel 会产生 panic
2. 重复关闭同一个 channel 会产生 panic
3. 向一个已关闭的 channel 发送消息会产生 panic
4. 从已关闭的 channel 读取消息不会产生 panic， 且能读出 channel 中还未被读取的消息，若消息均已被读取，则会读取到该类型的零值。
5. 从已关闭的 channel 读取消息永远不会阻塞，并且会返回一个为 false 的值，用以判断该 channel 是否已关闭 （x, ok := <- ch）
6. 关闭 channel 会产生一个广播机制，所有向 channel 读取消息的 goroutine 都会收到消息
7. channel 在 Golang 中是一等公民，他是线程安全的，面对并发问题，应首先想到 channel。
*/
func init() {
	// 双向通道
	pipelined := make(chan int, 10)

	fmt.Printf("管道可缓冲 %d 个数据 \n", cap(pipelined))

	pipelined <- 1
	fmt.Printf("管道中当前有 %d 个数据 \n", len(pipelined))

	go func() {
		fmt.Println("准备发送数据: 100")
		pipelined <- 100
	}()

	go func() {
		num := <-pipelined
		fmt.Printf("接收到的数据是: %d \n", num)
	}()
	chanType()
	chanRange()
	chanLock()
}

// Sender 定义只写信道类型
type Sender = chan<- int

// Receiver 定义只读信道类型
type Receiver = <-chan int

// 单向信道 只读与只写
func chanType() {
	var pipelined = make(chan int)

	go func() {
		var sender Sender = pipelined
		fmt.Println("准备发送数据:99")
		sender <- 99
	}()

	go func() {
		var receiver Receiver = pipelined
		num := <-receiver
		fmt.Printf("接收到的数据是: %d \n", num)
	}()

	time.Sleep(1)
}

// 遍历信道
func chanRange() {
	pipelined := make(chan int, 10)

	go func() {
		n := cap(pipelined)
		x, y := 1, 1
		for i := 0; i < n; i++ {
			pipelined <- x
			x, y = y, x+y
		}
		// 记得 close信道，不然会阻塞而不是结束
		close(pipelined)
	}()

	for k := range pipelined {
		fmt.Println(k)
	}
}

func increment(ch chan bool, x *int) {
	ch <- true
	*x = *x + 1
	<-ch
}

// 管道锁
func chanLock() {
	// 注意要设置容量为 1 的缓冲信道
	pipelined := make(chan bool, 1)

	var x int
	for i := 0; i < 1000; i++ {
		go increment(pipelined, &x)
	}

	// 确保所有的协程都已完成
	time.Sleep(10)
	fmt.Println("x 的值:", x)
}
