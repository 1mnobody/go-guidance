package basic

import (
	"fmt"
	"math"
	"runtime"
)

// Go 只有一种循环结构: for
func Loop()  {
	sum := 0
	// 初始化语句; 条件表达式; 后置语句
	for i := 0; i < 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 省略初始化和后置语句，这种情况下，可以去掉分号（相当于其他语言中的while语句）
	sum2 := 1
	for ; sum2 < 100; {
		sum2 += sum2
	}
	fmt.Println(sum2)

	sum3 := 3
	// 死循环
	for {
		if sum3 > 100 {
			fmt.Printf("break when sum is %v \n", sum3)
			break
		}
		sum3 += sum3
	}

	// if语句也可以在条件表达式前执行一个简单语句
	if i := 100; i < Sum(50, 60) {
		fmt.Println("smaller")
	} else {
		fmt.Println(i)
	}
}

func Sum(x, y int) int {
	return x + y
}


func Sqrt(x float64) {
	// 牛顿法 计算一个数的平方根
	z := x/2
	for x - z*z > 0.0001 || x - z*z < -0.0001 {
		z -= (z*z - x) / (2*z)
	}
	// 与 math 库计算出来的数据比较
	fmt.Println(z, "   ", math.Sqrt(x))
}


func Switch() {
	// 运行此函数，只会输出Windows，说明Go语言中的switch只会运行选定的case语句。
	// 当以fallthrough结束case时，运行完选定的case之后会再进入到后面的case继续执行。
	// 与其他语言一样，case条件的查找是自上而下的，找到一个匹配的值之后执行该case后面的语句
	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("Windows")
		//fallthrough   // 加入了 fallthrough，除了打印Windows，还会打印 OS X. 由于后面的case没有 fallthrough，执行完之后不再执行其他case语句
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println(os)
	}
}

// defer 语句会将函数推迟到外层函数返回之后执行，推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用。
// 这里Defer()函数返回后，Defer()内部的fmt.Println("defer 指定的函数") 才会被执行
func Defer() {
	// 推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。
	defer fmt.Println("defer 指定的函数1")
	defer fmt.Println("defer 指定的函数2")
	defer fmt.Println("defer 指定的函数3")

	fmt.Println("hello")
}