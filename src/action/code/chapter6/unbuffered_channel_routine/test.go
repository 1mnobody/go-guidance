package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// 一个单位的缓冲区，数据先放在缓冲区内，再被当前的goroutine读取（无缓冲的channel会导致此示例报错）
	c := make(chan int, 1)
	for {
		c <- 1
		res := <-c
		fmt.Println(res)

		random := rand.Intn(10)
		if random > 8 {
			close(c)
			_, ok := <-c
			fmt.Println(ok)
			return
		}
	}
}
