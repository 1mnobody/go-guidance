/**
每个Go程序都是由package 组成，程序从main包下的main方法开始运行，同一个文件夹下的所有go源文件，只能存在一个package（package名称一样）
*/
package main

/**
import 用于导入包，可以用圆括号导入多个包，一个包一行
*/
import (
	"./basic"
	"./go_routine"
	"./method_interface"
	"fmt"
	"time"
)

func init() {
	fmt.Println("init 会在 main之前执行")
}

func main() {
	//fmt.Println("Hello World")
	runBasic()
	runMethod_Interface()
	goroutine()
}

func runBasic() {
	fmt.Println(basic.ExportedVar)
	// unExportedVar为小写字母开头，无法访问此变量
	//fmt.Println(basic.unExportedVar)

	fmt.Println("-------- function --------")
	basic.Println("这是一个无返回值的函数")
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
	f1 := basic.Adder() // Adder返回一个函数（闭包），此函数接收一个int类型的参数
	f2 := basic.Adder()
	fmt.Println(f1(10), "  ", f2(10))
}

func runMethod_Interface() {
	fmt.Println("----------- 方法接收者 -----------")
	v := method_interface.Vertex{3, 4}
	i := method_interface.Number{3}
	// Abs方法中修改了X,Y的值，但是在打印后可以看到v中的X,Y没有变化，说明传入的对象v的一个副本；
	// 与此相对的有  接收者指针，此时在方法中修改接收者参数的值会导致原实例的值发生改变，详情见下面的v.Abs()与v.Scale()方法
	fmt.Println(v.Abs(), "--", i.Abs())
	fmt.Println(v.X, "==", v.Y)
	v.Scale(2.0)
	fmt.Println(v.X, "==", v.Y)

	fmt.Println("----------- 接收者既可以为值，也可以为指针，编译器会自动对其处理，转成目标接收者 -----------")
	// Abs()方法定义时以值为接收者，这里用Vertex指针作为接收者去执行方法，编译器会自动转换。相当于(*p).Abs()
	// 同样的，Scala方法的接收者是指针，上面的 v.Scale(2.0) 会被解释成 (&v).Scale(2.0)
	p := &method_interface.Vertex{6, 8}
	fmt.Println(p.Abs())

	fmt.Println("----------- 接口：一组方法的集合，接口变量可以保存实现了这些方法的具体实例 -----------")
	var a method_interface.Abser
	myInteger := method_interface.MyInteger(10)
	a = myInteger
	fmt.Printf("接收类型MyInteger：%T, %d\n", a, a.Abs())
	a = &myInteger
	fmt.Printf("接收类型MyInteger的地址：%T, %d\n", a, a.Abs())
	iv := method_interface.V{3, 4}

	// a = iv  // 报错，原因是为V定义的Abs()方法是以*V（V的指针）作为接收者
	// fmt.Println("接收类型V：", a, a.Abs())

	a = &iv
	fmt.Printf("接收类型V的地址 %T, %d\n", a, a.Abs())

	var i1 method_interface.I // **i1此时为nil接口**
	var t *method_interface.S
	fmt.Printf("接口中没有值，也没有具体类型 ，%v , %T\n", i1, i1)
	fmt.Println("nil接口无法进行方法调用，即当前调用i1.Method()会报错")
	fmt.Println("----------- 接收者为nil，方法也会被nil接收者调用 -----------")
	i1 = t // ** i1为S类型，但是其值为nil **
	i1.Method()

	i1 = &method_interface.S{"hahaha"}
	i1.Method()

	// 指定了零个方法的接口被称为 空接口
	fmt.Println("----------- 空接口可以保存任意类型的值 -----------")
	var itf interface{}
	itf = 100
	// 空接口可以用来处理未知类型的参数值，fmt.Println()方法就是接收 interface{} 类型的参数
	fmt.Printf("%T, %v\n", itf, itf)

	// 断言类型，类似于java中的 instanceof
	// 断言类型有两种返回方式，单值返回时如果不是对应的类型，则会报错。
	// 下面的示例用于判断itf是否是int类型，如果是，将该值返回，否则返回零值
	as1 := itf.(int)
	fmt.Println(as1)
	// as2 := itf.(string) // 报错 panic
	as3, ok := itf.(float64)
	fmt.Println(as3, "  ", ok)

	// type switches，根据类型进入不同switch分支
	// 类型选择与类型断言 i.(T) 的语法一致，只是具体的类型T变成了关键字type，具体查看下面的Do方法
	fmt.Println("----------- 类型选择 -----------")
	method_interface.Do(100)
	method_interface.Do("hahah")

	fmt.Println("----------- 定义结构体的String方法 -----------")
	person := method_interface.Person{"wuzhsh", 27}
	fmt.Println(person)

	fmt.Println("----------- error -----------")
	method_interface.ErrorDemo()
	e := method_interface.Run()
	if e != nil {
		fmt.Println(e)
	}

	fmt.Println("----------- reader -----------")
	method_interface.ReaderDemo()
	_, err := method_interface.Devide(100, 0)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("----------- 返回值按值传递 -----------")
	person2 := method_interface.GetInstance()
	// 这里可以看到返回的对象地址与在方法中创建的对象地址不同，所以这里拿到的是方法中创建的对象的一个副本（一个新的对象）
	fmt.Printf("in caller:%p \n", &person2)

	person3 := method_interface.GetAddress()
	fmt.Printf("in caller:&person3=%p ,person3=%p\n", &person3, person3)
}

func goroutine() {
	fmt.Println("------------ go routine ------------")
	go go_routine.Say("world")
	go_routine.Say("hello")

	fmt.Println("------------ 信道 ------------")
	// 创建一个信道，信道上的发送和接收操作在另一端准备好之前都会阻塞，使用信道前必须先创建信道
	//ch := make(chan int)
	go_routine.Sum_Invocation()
	go_routine.Range_Close()
	go_routine.Select()
	go_routine.MutexDemo()
	// 主线程休眠5s等待 所有的goroutine执行完
	time.Sleep(5 * time.Second)
}

func add(x, y int) int {
	return x + y
}
