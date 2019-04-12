/**
每个Go程序都是由package 组成，程序从main包下的main方法开始运行，同一个文件夹下的所有go源文件，只能存在一个package（package名称一样）
*/
package main

/**
import 用于导入包，可以用圆括号导入多个包，一个包一行
*/
import (
	"./basic"
	"./method_interface"
	"fmt"
)

func init() {
	fmt.Println("init 会在 main之前执行")
}

func main() {
	fmt.Println("Hello World")
	runBasic()
	//runMethod_Interface()
}

func runBasic() {
	fmt.Println(basic.ExportedVar)
	// unExportedVar为小写字母开头，无法访问此变量
	//fmt.Println(basic.unExportedVar)

	fmt.Println("-------- function --------")
	fmt.Println(basic.Add(1, 2))
	fmt.Println(basic.Swap("123", "abc"))
	fmt.Println(basic.NamedRet(12))
	fmt.Println("使用一个函数作为另外一个函数的参数：", basic.FuncParam(add, 1, 2))

	fmt.Println("-------- var --------")
	basic.Print()
	basic.BasicType()
	basic.TypeConversion()

	fmt.Println("-------- flow control --------")
	basic.Loop()
	basic.Sqrt(3)
	basic.Switch()
	basic.Defer()

	fmt.Println("-------- pointer --------")
	basic.PointerPrint()
	basic.PrintStructA()

	fmt.Println("-------- array --------")
	basic.ArrayDemo()
	basic.MakeDemo()

	fmt.Println("-------- map --------")
	basic.MapDemo()

	fmt.Println("-------- closure --------")
	f := basic.Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Print(f(), " ,")
	}
	fmt.Println()
	f1 := basic.Adder() //Adder返回一个函数，此函数接收一个int类型的参数
	f2 := basic.Adder()
	fmt.Println(f1(10), "  ", f2(10))
}

func runMethod_Interface() {
	v := method_interface.Vertex{3, 4}
	i := method_interface.Number{3}
	fmt.Println(v.Abs(), "   ", i.Abs())
}

func add(x, y int) int {
	return x + y
}
