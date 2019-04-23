package method_interface

import "fmt"

func GetInstance() Person {
	p := Person{"aaa", 12}
	fmt.Printf("in function: %p\n", &p)
	return p
}

func GetAddress() *Person {
	p := &Person{"bbb", 22}
	fmt.Printf("in function: &p=%p p=%p\n", &p, p)
	return p
}
