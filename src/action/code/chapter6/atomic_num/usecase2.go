package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	shutdown int64
	wg2      sync.WaitGroup
)

func main() {

	wg2.Add(2)
	go doWork("A")
	go doWork("B")

	time.Sleep(2 * time.Second)
	fmt.Println("Shutdown now")
	// Set shutdown value safely
	// atomic 的 Store\Load 方法可以在多个routine之间安全的访问数值类型的值
	atomic.StoreInt64(&shutdown, 1)

	wg2.Wait()
}

func doWork(name string) {
	defer wg2.Done()
	for {
		fmt.Printf("Doing %s Work\n", name)
		time.Sleep(200 * time.Millisecond)

		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s Down\n", name)
			break
		}
	}
}
