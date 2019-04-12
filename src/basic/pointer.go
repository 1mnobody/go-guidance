package basic

import "fmt"

// 指针的声明，初始值为nil
var p *int

func PointerPrint() {
	fmt.Println(p)
	i := 100
	// & 操作符会生成一个指向其操作数的指针
	p = &i
	// * 操作符解指针，表示指针指向的值
	fmt.Println(p, "   ", *p)
	*p = 20
	fmt.Println(i)

}