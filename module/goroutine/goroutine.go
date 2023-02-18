package goroutine

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

func init_goroutine() {
	log.Info("module goroutine")
	setup()
}

func portal1(channel chan string) {
	time.Sleep(time.Second)
	channel <- "portal1"
}

func portal2(channel chan string) {
	channel <- "portal2"
}

func setup() {
	R1 := make(chan string)
	R2 := make(chan string)

	go portal1(R1)
	go portal2(R2)

	select {
	case op1 := <-R1:
		fmt.Println(op1)
	case op2 := <-R2:
		fmt.Println(op2)

	}
}
