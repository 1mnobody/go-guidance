package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var b bytes.Buffer
	// 通过 string 构造一个byte数组，将数据写到Buffer中
	b.Write([]byte("Hello"))
	// 拼接一个字符串到Buffer中
	fmt.Fprint(&b, " World!")
	// 将数据写到标准输出
	io.Copy(os.Stdout, &b)
}
