package method_interface

import "fmt"

func Do(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Printf("int , value: %d\n", i)
	case string:
		fmt.Printf("string, value: %s\n", i)
	case byte:
		fmt.Printf("byte, value: %d", i)
	}
}
