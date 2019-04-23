package method_interface

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

// 通常函数会返回一个error值，调用此函数的代码应该判断error值是否为nil来进行错误处理
func ErrorDemo() {
	s, error := strconv.Atoi("abc")
	if error == nil {
		fmt.Println(s)
	} else {
		fmt.Printf("convert error, %v\n", error)
	}
}

func Devide(i1, i2 int64) (int64, error) {
	if i2 == 0 {
		return 0, errors.New("被除数不能为0")
	}
	return i1 / i2, nil
}

// 自定义error类型，实现了  Error() 方法
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

// error类型是go的内建接口
func Run() error {
	fmt.Println("run ...")
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}
