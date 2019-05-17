package basic

import "fmt"

var m map[int]string

// map的定义
// map[key 类型]value类型 {
//     key: value 的类型{ // 类型可省略
//         value
//     },   // 逗号是必须的
// }
var m2 = map[int]string{
	1: "aa",
	2: "bb",
	3: "cc",
}

var m3 = map[string]A{
	"a": A{
		1, 1,
	},
	"b": A{
		2, 2,
	},
}

var m4 = map[string]A{
	"a": {2, 2},
	"b": {3, 3},
}

func MapDemo() {
	// 映射的零值为nil，不能添加键值对，以下语句会报错
	//m[0] = "haha"

	// 检测键是否存在
	value, ok := m2[1]
	if ok {
		fmt.Println(value)
	}

	m2[0] = ".."
	for k, v := range m2 {
		fmt.Print("[", k, ",", v, "] ")
	}
	fmt.Println()

	am := make(map[int]string)
	am[0] = "00"
	am[2] = "11"
	am[3] = "33"

	for k, v := range am {
		fmt.Print("[", k, ",", v, "] ")
	}
	fmt.Println()
}
