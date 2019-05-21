package main

import (
	"./entities"
	"fmt"
)

func main() {
	// Admin 中包含一个inner type： user，user中定义的都是exported field，可以直接通过outer type来访问这些field
	admin := entities.Admin{
		Right: 10,
	}

	admin.Name = "hahaha"
	admin.Email = "123@123.com"

	fmt.Printf("admin info: [%v] \n", admin)
}
