package main

import (
	"../../work"
	"log"
	"sync"
	"time"
)

var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

type namePrinter struct {
	name string
}

func (p *namePrinter) Task() {
	log.Println(p.name)
	time.Sleep(time.Second)
}

func main() {
	// pool 中会创建一个channel，channel的缓冲区为2，会启动两个goroutine，range over 这个channel，
	// Run方法会将work推到channel中，启动的两个goroutine会拿到work，执行其Task方法
	p := work.New(2)

	var wg sync.WaitGroup
	wg.Add(100 * len(names))

	for i := 0; i < 100; i++ {
		for _, name := range names {
			np := namePrinter{
				name,
			}
			go func() {
				// p.Run() 会将Worker推送到channel中，p中启动的两个goroutine会不断的
				// 获取Worker并执行其Task()方法
				p.Run(&np)
				wg.Done()
			}()
		}
	}
	wg.Wait()
	p.Shutdown()
}
