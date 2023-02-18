package goroutine

import (
	"fmt"
	"time"
)

/* time
1. Timer 单次定时
2. Ticker 周期定时
*/
func init() {
	// 1.timer 的简单使用
	//timer1 := time.NewTimer(2 * time.Second)
	//t1 := time.Now()
	//fmt.Printf("t1:%v\n", t1)
	//t2 := <- timer1.C
	//fmt.Printf("t2:%v\n", t2)

	// 2.验证timer只能相应1次
	//timer2 := time.NewTimer(time.Second)
	//for {
	//	<-timer2.C
	//	fmt.Println("时间到")
	//}

	// 3.timer实现延时功能
	//time.Sleep(time.Second)
	//timer3 := time.NewTimer(2 * time.Second)
	//fmt.Println("timer start")
	//<- timer3.C
	//fmt.Println("2秒到")
	//<- time.After(2*time.Second)
	//fmt.Println("2秒到")

	// 4.停止计时器
	//timer4 := time.NewTimer(2 * time.Second)
	//go func() {
	//	<- timer4.C
	//	fmt.Println("定时器执行了")
	//}()
	//b := timer4.Stop()
	//if b {
	//	fmt.Println("timer4 已经关闭")
	//}

	// 5.重置定时器
	//timer5 := time.NewTimer(2 * time.Second)
	//timer5.Reset(5 * time.Second)
	//fmt.Println(time.Now(), "line1")
	//fmt.Println(<-timer5.C, "line2")
	//
	//for  {
	//
	//}

	// Ticker
	ticker := time.NewTicker(1 * time.Second)
	i := 0
	// 子协程
	go func() {
		for {
			<-ticker.C
			i++
			fmt.Println(<-ticker.C)
			if i == 5 {
				ticker.Stop()
			}
		}
	}()
	for {

	}
}
