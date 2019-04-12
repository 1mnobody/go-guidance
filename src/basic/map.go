package basic

import "fmt"

var m map[int]string
var m2 = map[int]string{
	1: "aa",
	2: "bb",
	3: "cc",
}

func MapDemo()  {
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
