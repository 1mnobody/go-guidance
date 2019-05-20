package main

// matchers 包中包含了要注册的matcher（使用init方法进行注册，所以这里要import，但是main.go 的代码中未使用此包，
// 故使用下划线的方式import）
import (
	_ "./matchers"
	"./search"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
