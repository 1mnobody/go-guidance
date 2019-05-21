package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	// 多个routine同时访问counter
	counter int
	wg      sync.WaitGroup
)

func main() {
	wg.Add(2)

	go incCounter()
	go incCounter()

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incCounter() {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		value := counter
		// Gosched() 方法会让 routine 让出处理器，允许其他的goroutine执行
		runtime.Gosched()
		value++
		counter = value
	}
}
