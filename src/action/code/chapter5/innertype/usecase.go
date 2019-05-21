package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

type admin struct {
	// user在这里作为embedded type，embedded type 的方法可以用 embedding type 去调用（user的方法可以被admin直接调用）
	user
	level string
}

/**
  特别注意：如果admin以以下的方式定义，则无法直接调用 user的方法
type admin struct {
	u user
	level string
}
*/

func main() {
	ad := admin{
		user{"ad", "123@123.com"},
		"super",
	}

	//
	ad.user.notify()
	ad.notify()
}
