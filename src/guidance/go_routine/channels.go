package go_routine

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	// 将结果传入channel
	c <- sum
}

func Sum_Invocation() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	c := make(chan int)
	// 带缓冲的信道，仅当缓冲区满时，向其发送数据才会阻塞；而接收数据会在缓冲区为空时阻塞
	// c := make(chan int, 100)
	go sum(s[:len(s)/2], c) //[0,4)
	go sum(s[len(s)/2:], c) //[4,9)
	//x, y := <-c, <-c        // 从信道中接收数据
	x, ok1 := <-c // 这种格式获取信道中的数据，第二个值用于表示信道是否被关闭
	y, ok2 := <-c
	fmt.Println(ok1, ok2, "  ", x, "  ", y)
}

func Range_Close() {
	c := make(chan int)
	go produce(c)
	// range会不停的获取信道c中的值，直到发送者关闭信道，
	// range over一个channel时，只允许使用一个迭代变量（从channel中获取数据时可以接收到两个值：v, ok <-ch）
	for x := range c {
		fmt.Printf("%v  ", x)
	}
	fmt.Println()
}

func produce(c chan int) {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		c <- i
	}
	// 关闭信道时，接收方会跳出循环
	// 注意：只有发送者才能关闭信道，而接收者不能。向一个已经关闭的信道发送数据会引发程序恐慌（panic）。
	close(c)
}

func Select() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%v ", <-c)
		}
		// 这里,quit在for循环执行完之前，一直是没有数据进入的，所以在下面的select语句中会先执行 第一个case或者default
		quit <- 0
	}() // 定义了一个匿名函数，在此函数中向两个信道中写数据

	x, y := 0, 1
	for {
		// select语句可以让goroutine等待多个通信操作，select会阻塞到某个分支可以继续执行时执行该分支，如果多个分支都可以执行
		// 则会随机选择一个分支，如果有default分支，则在其他分支中的信道数据没有准备好时就会执行default分支。
		// close掉的channel，也是可以执行的，尝试Select2方法
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		default:
			fmt.Println("执行default，休眠10ms")
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func Select2() {
	c := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		close(c)
	}()

	var count = 0
	for {
		select {
		case <-c:
			if count > 5 {
				return
			}
			count++
			fmt.Println("chan closed" + string(count))
		default:
			fmt.Println("---------")
		}

	}
}
