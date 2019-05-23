package main

import "log"

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	log.Println("message")
	// Fatalln() 等价于先调用Println()，再调用os.Exit(1)，Fatalln()后面的语句不会再执行
	//log.Fatalln("fatal message")
	// Panicln(msg) 等价于先调用Println(mss)，再调用panic(msg)
	log.Panicln("panic message")
}
