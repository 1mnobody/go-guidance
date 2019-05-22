package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	sendToClosedChannel()
	recvFromClosedChannel()
}

func sendToClosedChannel() {
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan string, 10)
	go func() {
		v := <-ch
		fmt.Println(v)
		wg.Done()
	}()

	// 注意，往一个close掉的channel（无论是buffered还是unbuffered）中发送数据会造成 panic: send on closed channel
	go func() {
		ch <- "ahahahaha"
		wg.Done()
	}()
	close(ch)
	wg.Wait()
}

func recvFromClosedChannel() {
	var wg sync.WaitGroup
	wg.Add(2)
	// 存在缓冲区的channel，发送数据后马上close，在接收端也可以接收到数据。
	ch := make(chan string, 10)

	go func() {
		for i := 0; i < 10; i++ {
			msg := "msg" + strconv.Itoa(i)
			ch <- msg
		}
		wg.Done()
		fmt.Println("Finish sending")
		close(ch)
	}()

	go func() {
		// 等待5s，让数据发送到channel中
		time.Sleep(5 * time.Second)
		for i := 0; i < 15; i++ {
			// 接收两个值，第二个值为bool类型，当能够从channel中接收到数据时，此值为true，
			// 当 ** channel关闭并且channel中没有数据了 **，此值为false
			v, ok := <-ch

			fmt.Printf("[value %s, ok %t, channel closed? %t]\n", v, ok, ch.clo)
		}
		wg.Done()
	}()

	wg.Wait()
}
