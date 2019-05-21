package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
	mutex   sync.Mutex
)

func main() {
	wg.Add(2)

	go incCounter()
	go incCounter()

	wg.Wait()
	fmt.Printf("Final Counter:%d \n", counter)
}

func incCounter() {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// 一次只允许一个routine进入这段代码
		mutex.Lock()
		value := counter
		runtime.Gosched()
		value++
		counter = value
		mutex.Unlock()
	}
}
