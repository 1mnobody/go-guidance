package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// 分配一个逻辑处理器供routine调度器使用
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	// wg.Add(2) 对应两个goroutine, 每个routine中都会执行wg.Done()，Done方法执行两次后，执行wg.Wait()的线程会被唤醒
	wg.Add(2)
	fmt.Println("执行goroutine")

	// 只分配了一个逻辑处理器，所以下面的两个goroutine实际上是串行执行，要么先输出完大写的字母，要么先输出完小写的字母。
	// 可将 runtime.GOMAXPROCS(1) 中的 1 改为 2 作比较
	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for ch := 'a'; ch < 'a'+26; ch++ {
				fmt.Printf("%c", ch)
			}
			fmt.Println()
		}
	}()

	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for ch := 'A'; ch < 'A'+26; ch++ {
				fmt.Printf("%c", ch)
			}
			fmt.Println()
		}
	}()

	fmt.Println("等待程序执行完成..")
	wg.Wait()

	fmt.Println("程序执行完成")

next:
	for outer := 2; outer < 100; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next // continue外层循环，假设outer=4， outer % 2 == 0， continue外层循环，则 outer=5 再执行循环
			}
		}
		fmt.Printf("%s:%d\n", "a", outer)
	}
	fmt.Println("Completed", "a")
}
