package method_interface

import (
	"fmt"
	"io"
	"strings"
)

func ReaderDemo() {
	r := strings.NewReader("abcdefg hijklmn opq")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		if err == io.EOF {
			break
		}
		fmt.Printf("n = %v err = %v b = %c\n", n, err, b)
	}
}
