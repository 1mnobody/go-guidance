package basic

// 接收两个参数，一个返回值，大写开头，所以可以直接在(basic)包外使用
func Add(x int, y int) int {
	return x + y
}

// 连续的两个或多个命名形参类型相同，可以省略前面的参数类型
func Add2(x, y int) int {
	return x + y
}

// 多值返回
func Swap(x, y string) (string, string) {
	return y, x
}

// 返回值可以被命名，它们会被视作定义在函数顶部的变量。下面的例子相当于在外部定义了x, y两个变量
func NamedRet(sum int) (x, y int) {
	x = sum / 2
	y = sum + 1
	return
}

// 此函数接收一个函数作为参数
func FuncParam(function func(x, y int) int, param1, param2 int) int {
	return function(param1, param2)
}
