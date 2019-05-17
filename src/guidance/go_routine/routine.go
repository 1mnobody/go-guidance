package go_routine

import (
	"fmt"
	"time"
)

// goroutine是由Go运行时管理的轻量级线程
// go func(x,y,z) 会启动一个goroutine并执行 func(x,y,z)，func,x,y,z的求值是在当前的goroutine中，而func的执行则是在
// 新的goroutine中

func Say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println(s)
	}
}
