package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	// court是一个无缓冲的channel,所以 ** 发送会阻塞直到数据被接收，接收也会阻塞直到能读到数据 **
	court := make(chan int)
	wg.Add(2)

	go player("Nadal", court)
	go player("Djokovic", court)

	court <- 1

	wg.Wait()
}

func player(name string, court chan int) {
	defer wg.Done()

	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("Player %s won \n", name)
			return
		}

		n := rand.Intn(100)

		if n%13 == 0 {
			fmt.Printf("Player %s missed\n", name)
			close(court)
			return
		}
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		court <- ball
	}
}
