package method_interface

import "math"

// Go没有类，有结构体类型，结构体类型可以定义方法（注意方法与函数的区别：方法是带接收者的函数，接收者的意思见下面的Abs方法）
type Vertex struct {
	// 注意，变量如果是小写，则在包的外部无法读取变量
	//x, y float64
	X, Y float64
}

type Number struct {
	I int
}

// 接收者 类型的定义和方法定义必须在同一个包内
// func 关键字 与方法名之间的列表就指定了方法接收者，此例中，方法接收者为Vertex类型
func (v Vertex) Abs() float64 {
	res := math.Abs(math.Sqrt(v.X*v.X + v.Y*v.Y))
	v.X = 100
	v.Y = 200
	return res
}

// 接收者为指针：可以修改接收者指向的值
// 使用指针作为接收者的原因：
// 1、可以修改接收者指向的值；
// 2、避免在每次调用时创建副本。
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (i Number) Abs() int {
	if i.I < 0 {
		return -i.I
	} else {
		return i.I
	}
}
