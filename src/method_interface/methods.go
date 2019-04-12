package method_interface

import "math"

// Go没有类，有结构体类型，结构体类型可以定义方法（注意方法与函数的区别：方法是带接收者的函数，接收者的意思见下面的Abs方法）
type Vertex struct {
	// 注意，变量如果是小写，则在包的外部无法读取变量
	//x, y float64
	X,Y float64
}

type Number struct {
	I int
}

// func 关键字 与方法名之间的列表就指定了方法接收者，此例中，方法接收者为Vertex类型
func (v Vertex) Abs() float64 {
	return math.Abs(math.Sqrt(v.X*v.X + v.Y*v.Y))
}

func (i Number) Abs() int {
	if i.I < 0 {
		return -i.I
	} else {
		return i.I
	}
}