package main

import "fmt"

func main() {
	// call() 方法中从panic中recover，main函数无感知
	call()
	fmt.Println("main执行完毕")
}

func call() {
	// defer要定义在可能出现panic的方法前，否则程序走不到defer语句就因为panic退出了
	defer func() {
		// recover() 方法只有定义在defer中才是有效的
		if err := recover(); err != nil {
			fmt.Println("出现了panic, 通过recover() 获取到panic：", err)
		} else {
			fmt.Println("收尾工作")
		}
	}()
	panicFunc()
	fmt.Println("call执行完毕")
}

func panicFunc() {
	panic("encounter unknown error")
}
