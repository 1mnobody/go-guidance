package basic

import (
	"fmt"
	"math/cmplx"
)

// var 语句可以出现在包或函数级别，没有初始值的变量声明会被赋予 零值：bool -> false，数值类型 -> 0，string -> ""
var c, py, java string

// const关键字定义常量，常量可以是字符、字符串、布尔值或数值。常量不能用 := 语法声明。
const name = "abc"
const (
	conA = 1
	conB = "1"
	conC = true
)

var (
	b   bool       = false
	max uint32     = 1<<32 - 1
	z   complex128 = cmplx.Sqrt(-5 + 12i)
)

func Print() {
	// 会覆盖包级别的变量
	var py, java = "222", "123"
	fmt.Println(c, py, java)

	// 简洁赋值语句 := 可在类型明确的地方代替 var 声明，Go的类型推导会自动通过右值得出变量是什么类型的。对于数值常量，变量的类型取决于数值的精度
	// 注意：函数外的每个语句都必须以关键字开始（var, func等），因此 := 结构不能在函数外使用
	i1, i2 := 11, 22
	fmt.Printf("%T, %T, value: %v, %v \n", i1, i2, i1, i2)
	f := 0.123
	fmt.Printf("%T \n", f)
}

func BasicType() {
	fmt.Printf("Type: %T Value: %v\n", b, b)
	fmt.Printf("Type: %T Value: %v\n", max, max)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

// 类型转换，T(v)会将值v转换成类型T，Go 在不同类型的项之间赋值时需要显式转换
func TypeConversion() {
	var i int = 43
	var f float32 = float32(i)
	var u uint = uint(f)

	si := 43
	sf := float32(si)
	su := uint(sf)

	fmt.Println(i, f, u)
	fmt.Println(si, sf, su)
}
