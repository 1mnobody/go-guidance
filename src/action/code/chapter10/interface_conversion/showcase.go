package main

import (
	"fmt"
)

// 编译时的隐式接口转换

type Mover interface {
	Move()
}

type Locker interface {
	Lock()
	Unlock()
}

type MoveLocker interface {
	Mover
	Locker
}

type bike struct {
}

func (b bike) Move() {
	fmt.Println("moving the bike")
}

func (b bike) Lock() {
	fmt.Println("locking the bike")
}

func (b bike) Unlock() {
	fmt.Println("unlocking the bike")
}

func main() {
	var ml MoveLocker
	var m Mover

	ml = bike{}

	// Mover 与 MoveLocker 都定义了  move 方法，所以可以将 MoveLocker 实例隐式转换成 Mover 实例
	m = ml

	// Cannot use 'm' (type Mover) as type MoveLocker in assignment Type does not implement 'MoveLocker'
	// as some methods are missing: Lock() Unlock()
	// ml = m
	// 接口m没有定义Lock() Unlock()方法，编译器无法完成隐式转换，但是可以通过类型断言完成类型转换
	// 此时，ml中存放的值是 m 的一个拷贝，即 类型断言 会创建一个新的实例，其数据与原始实例一致
	fmt.Printf("%p\n", &ml)
	ml = m.(bike)
	fmt.Printf("%p, %p", &ml, &m)
}
