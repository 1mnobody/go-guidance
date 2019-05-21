package main

import "fmt"

// embedded type 实现了接口，embedding type 也可被认为“实现”了接口
type user struct {
	name  string
	email string
}

type notifier interface {
	notify()
}

func (u *user) notify() {
	fmt.Printf("user: [" + u.name + ", " + u.email + "]")
}

type admin struct {
	user
	level string
}

func (a *admin) notify() {
	fmt.Println("admin:[" + a.name + ", " + a.email + "]")
}

func main() {
	// 这里创建一个admin， 实现接口的是user，传入admin的地址，可以实现notify()的调用。
	// ** 增加一个admin的实现，可以发现如果admin本身也实现了接口，则调用的是admin本身的实现，同时注意，输出的a.name，
	// a.email 其实是 admin 中的 user 的属性 **
	a := admin{
		user{"zhangsan", "zhangsan@123.com"},
		"aaaa",
	}
	sendNotification(&a)
}

func sendNotification(n notifier) {
	n.notify()
}
