package goroutine

import (
	"fmt"
	"time"
)

/*
惰性生成器
*/
var resume chan int

func integers() chan int {
	yield := make(chan int)
	count := 0
	go func() {
		for {
			yield <- count
			count++
		}
	}()
	return yield
}

func generateInteger() int {
	return <-resume
}

func init() {
	resume = integers()
	fmt.Println(generateInteger())
	fmt.Println(generateInteger())
	fmt.Println(generateInteger())
	time.Sleep(5)
	fmt.Println(generateInteger())
	fmt.Println(generateInteger())

}
