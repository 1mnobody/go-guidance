package basic

import "fmt"

// 结构体就是一组字段的集合
type A struct {
	X int
	Y int
}

func PrintStructA() {
	a := A{1,2}
	a.X = 3
	fmt.Println(a.X, a.Y)

	// p 是一个指针，指向一个结构体A的实例a
	p := &a
	// 在Go中，允许使用隐式间接引用，编译器完成对指针的解引用，即p.Y (隐式间接引用）与(*p).Y是一样的效果
	fmt.Printf("(*p).Y: %v  -  p.Y: %v  \n",(*p).Y, p.Y)

	// 使用var关键字来创建结构体
	var (
		a1 = A{1,2}	// 创建一个结构体，指定其字段的值
		a2 = A{}		  	// 创建一个结构体，字段值初始化为零值
		a3 = A{X:1}		  	// 创建一个结构体，指定部分字段的值，其余字段被初始化为零值
		pa = &A{1,2}	// 使用&作为前缀，会返回一个指向结构体的指针
	)
	fmt.Printf("a1.X: %v, a2.X: %v, a3.X: %v, a4.X: %v \n", a1.X, a2.X, a3.X, pa.X)
	fmt.Printf("%v -- %v \n", a1, p)
}