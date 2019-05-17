package method_interface

import "math"

// 接口是由一组方法签名定义的集合，接口类型的变量可以保存任何实现了这些方法的值，即，类型通过实现接口的所有方法来实现这个接口
type Abser interface {
	Abs() int32
}

// 这种类型定义类似于C中的typedef
type MyInteger int32

func (i MyInteger) Abs() int32 {
	if i < 0 {
		return int32(-i)
	}
	return int32(i)
}

type V struct {
	X, Y int32
}

func (v *V) Abs() int32 {
	return int32(math.Sqrt(float64(v.X*v.X) + float64(v.Y*v.Y)))
}
