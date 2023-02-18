package goroutine

import (
	"fmt"
	"runtime"
	"time"
)

func init() {
	//sched()
	//exit()
	maxProc()
}

// runtime.Gosched() 让出CPU时间片，重新等待安排任务
func sched() {
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("world")
	// 主协程
	for i := 0; i < 2; i++ {
		// 切一下，再次分配任务
		runtime.Gosched()
		fmt.Println("hello")
	}
}

// runtime.Goexit() 退出当前协程
func exit() {
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			// 结束协程
			runtime.Goexit()
			defer fmt.Println("C.defer")
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()
}

func a() {
	for i := 0; i < 10; i++ {
		fmt.Println("A.", i)
	}
}

func b() {
	for i := 0; i < 10; i++ {
		fmt.Println("B.", i)
	}
}

func maxProc() {
	runtime.GOMAXPROCS(2)
	go a()
	go b()
	time.Sleep(time.Second)
}
