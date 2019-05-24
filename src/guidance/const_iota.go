package main

import "fmt"

// 常量定义，当后面的变量没有赋值时，默认与前一个变量一样的值
const (
	a = 1
	b
	c = 10
	d
)

// iota 表示一个整数序列，声明的常量依次递增，如下面的示例，e=0,f=1,g=2,h=3
const (
	e = iota
	f
	g
	h
) // iota在每个const定义中，都被初始化为0

const ( // 这里，iota又被初始化为0，且多次使用不会影响其递增效果
	i = iota
	j
	k = iota
	l
)

const ( // iota的值受const中定义的常量数而递增
	// 这里定义了5个变量，p为第4个变量，iota此时为3（从0递增）
	m = 5
	n = 19
	o = 100
	p = iota
	q
)

/**
1 1 10 10
0 1 2 3
0 1 2 3
5 19 100 3 4
*/
func PrintConst() {
	fmt.Println(a, b, c, d)
	fmt.Println(e, f, g, h)
	fmt.Println(i, j, k, l)
	fmt.Println(m, n, o, p, q)
}
