package goroutine

import (
	"fmt"
	"strconv"
	"sync"
)

func initSync() {
	demoSyncMap()
}

var x int64
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	wg.Done()
}
func demoWaitGroup() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}

var m sync.Map

func demoSyncMap() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)
			value, _ := m.Load(key)
			fmt.Printf("k=:%v,v=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
