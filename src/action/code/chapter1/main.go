package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

func printer(ch chan int) {
	for i := range ch {
		fmt.Println(Goid(), " -- Received ", i)
	}
	wg.Done()
}

func sleep(seconds int32) {
	time.Sleep(time.Duration(seconds) * time.Second)
	fmt.Println(Goid(), " -- in sleep")
	wg.Done()
}

func Goid() int {
	// 获取goroutine id
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic recover:panic info:%v", err)
		}
	}()

	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

func main() {
	c := make(chan int)
	go printer(c)
	wg.Add(2)
	for i := 0; i <= 10; i++ {
		c <- i
	}
	go sleep(5)
	close(c)
	fmt.Println(Goid(), " -- Close chan")
	wg.Wait()
}
