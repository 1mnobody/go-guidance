package basic

import "fmt"

// 闭包，返回一个函数，此函数可以访问其函数体之外的变量，换句话说，函数与这些变量绑定在一起构成了闭包
func Adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		fmt.Printf("闭包中的 sum 变量地址： %p \n", &sum)
		return sum
	}
}

func Fibonacci() func() int {
	var i, j = 0, 1
	return func() int {
		res := i
		i, j = j, i+j
		return res
	}
}
