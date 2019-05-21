package main

import "fmt"

type notifier interface {
	notify(in string) string
}

type notifier2 interface {
	notify2()
}

type user struct {
	name  string
	email string
}

type admin struct {
	name  string
	email string
}

func (u *user) notify(in string) string {
	return "notify of [" + u.name + ", " + u.email + "] -- " + in
}

func (u user) notify2() {
	fmt.Println(u.name + " -- " + u.email)
}

func (a *admin) notify(in string) string {
	return "admin -- [" + a.name + "," + a.email + "]"
}

func main() {
	u := user{"zhangsan", "zhangsan@123.com"}
	// 注意这里的区别：如果直接调用方法，编译器会自动处理类型 调用会被编译器转换为 (&u).notify(..)
	u.notify("haha")
	// 接口的接收规则：
	// 1. 如果使用指针接收者来实现一个接口，那么只有指向那个类型的指针才实现对应的接口
	// 2. 如果使用值接收者来实现一个接口，那么那个类型的值和指针都能够实现对应的接口
	sendNotification(&u) // 指针接收者 只有指针才是实现了对应的接口

	// 值接收者，则对应类型的值以及指针都是实现了该接口
	sendNotification2(u)
	sendNotification2(&u)

	// 实现了notify方法的都可以作为参数传入
	a := &admin{"admin", "admin123@123.com"}
	sendNotification(a)
}

func sendNotification(n notifier) {
	res := n.notify("hello")
	fmt.Println(res)
}

func sendNotification2(n notifier2) {
	n.notify2()
}
