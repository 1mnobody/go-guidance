package method_interface

import "fmt"

type I interface {
	Method()
}

type S struct {
	C string
}

func (s *S) Method() {
	if s == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println("pass")
}
