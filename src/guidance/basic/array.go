package basic

import "fmt"

// 声明一个包含10个元素的int类型数组，数组不能改变大小
var a [10]int

func ArrayDemo() {
	a[0] = 1
	a[1] = 2
	a[3] = 3
	a[4] = 4

	fmt.Println(a)

	// 初始化数组
	array := [3]int{1, 2, 3}
	fmt.Println(array)

	// 数组的切片，半开区间，前闭后开的区间，下面的操作得到的是a[1],a[2],a[3] 3个元素
	// 注意：切片并不会创建新的空间，只是展示原数组的数据，修改切片的数据会导致原数组的数据被修改（其他包含相同元素的切片也可以
	// 观察到相应的修改）
	// len(s) 可以得到切片s的长度（即切片中的元素个数），cap(s)得到切片的容量，切片容量是指是从它的第一个元素开始数，
	// 到其底层数组元素末尾的个数。参考下面的sliceOfa，此切片从数组的第1个元素开始，所以其容量是9（数组长度为10）
	sliceOfa := a[1:4]
	slice2 := a[2:6]
	fmt.Printf("第一个切片，类型: %T, %v 长度: %v, 容量: %v , 另一个切片：%v\n",
		sliceOfa, sliceOfa, len(sliceOfa), cap(sliceOfa), slice2)
	sliceOfa[0] = 100
	sliceOfa[1] = 200
	fmt.Printf("修改切片的内容，导致原数组的数据被修改： %v 另一个切片观察到的值：%v \n", a, slice2)

	// 切片的方式创建一个数组，注意与数组创建时的区别（不指定长度），下面的过程是创建了一个数组，然后构建一个切片，此切片引用了该数组
	slince3 := []int{1, 2, 3, 4, 5}
	fmt.Println(slince3)

	// 切片上下界的默认值，下界默认值为0，上界默认值为数组长度
	defaultSlice0 := a[:]  // 0~9 (下界为0，数组a长度为10，所以上界为10）
	defaultSlice1 := a[:2] // 0~1 (上界是开区间，故不包括2）
	defaultSlice2 := a[3:] // 3~9
	fmt.Printf("slice0: %v, slice1: %v, slice2: %v \n", defaultSlice0, defaultSlice1, defaultSlice2)
}

// 内建函数make可以用来创建切片，可通过make来创建动态数组
func MakeDemo() {
	// make 函数会分配一个元素为零值的数组并返回一个引用了它的切片
	s := make([]int, 5)      // 创建一个长度为5的切片（底层数组长度为5）
	s2 := make([]int, 5, 10) // 创建一个长度为5，容量为10的切片（底层数组长度为10）
	fmt.Printf("切片s的长度 %v, 容量 %v    ", len(s), cap(s))
	fmt.Printf("切片s2的长度 %v, 容量 %v\n", len(s2), cap(s2))

	// 切片的切片
	sliceOfSlice := [][]string{
		[]string{"a", "b", "c"},
		[]string{"1", "2", "3"},
	}
	fmt.Println(sliceOfSlice)

	appendSlice := []int{1, 2, 3, 4, 5}
	fmt.Printf("%p ,length: %v, capacity: %v, content %v \n", appendSlice, len(appendSlice), cap(appendSlice), appendSlice)
	appendSlice = append(appendSlice, 6, 7, 8, 9, 0)
	fmt.Printf("after append, %p, length: %v, capacity: %v, content %v \n", appendSlice, len(appendSlice), cap(appendSlice), appendSlice)

	// for range遍历切片，range操作会返回两个值，第一个是元素下标，第二个是元素值，
	// 和py一样，如果不需要使用元素下标，可以用 _ 来忽略它
	for i, v := range appendSlice {
		fmt.Print("appendSlice[", i, "]=", v, "  ")
	}
	fmt.Println()
}
