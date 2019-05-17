package method_interface

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	//return "[Name:" + p.Name + " - Age:" + strconv.FormatInt(int64(p.Age), 10) + "]"
	return fmt.Sprintf("[Name:%v - Age:%v]", p.Name, p.Age)
}
