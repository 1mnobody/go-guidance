package main

import (
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int32
	wg      sync.WaitGroup
)

func main() {
	wg.Add(2)

	go incCounter()
	go incCounter()

	wg.Wait()
}

func incCounter() {
	defer wg.Done()

	// 使用atomic包在多个routine之间安全的访问int类型
	for count := 0; count < 2; count++ {
		runtime.Gosched()
		atomic.AddInt32(&counter, 1)
	}

}
