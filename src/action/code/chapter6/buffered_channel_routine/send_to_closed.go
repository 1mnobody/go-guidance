package main

import (
	"fmt"
	"sync"
)

func main() {
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
