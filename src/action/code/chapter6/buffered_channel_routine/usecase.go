package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutine = 4
	taskLoad        = 10
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	tasks := make(chan string, taskLoad)

	// 启动4个goroutine，tasks的缓冲数为10个
	wg.Add(numberGoroutine)
	for gr := 1; gr <= numberGoroutine; gr++ {
		go worker(tasks, gr)
	}

	// 10个task放到了tasks的缓冲区
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task %d", post)
	}

	fmt.Println("close tasks channel")
	close(tasks)
	wg.Wait()
}

func worker(tasks chan string, worker int) {
	defer wg.Done()

	for {
		task, ok := <-tasks

		if !ok {
			fmt.Printf("Worker: %d, Shutting Down %s \n", worker, task)
			return
		}

		fmt.Printf("Worker: %d , Start %s \n", worker, task)

		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Printf("Worker: %d, Complete: %s \n", worker, task)
	}
}
