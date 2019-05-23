package main

import (
	"../../pool"
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	maxGoroutines   = 25
	pooledResources = 2
)

type dbConnection struct {
	ID int32
}

var idCounter int32

func (c *dbConnection) Close() error {
	log.Println("Close: Connection ", c.ID)
	return nil
}

func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create New Connection", id)
	return &dbConnection{id}, nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	p, err := pool.New(createConnection, pooledResources)
	if err != nil {
		log.Println(err)
	}

	for query := 0; query < maxGoroutines; query++ {
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}

	wg.Wait()

	log.Println("Shutdown")
	p.Close()
}

func performQueries(q int, p *pool.Pool) {
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}

	defer p.Release(conn)

	time.Sleep(time.Duration(rand.Intn(1000)) * time.Microsecond)
	// conn.(*dbConnection) ： 类型断言，conn.(*dbConnection)断言conn为*dbConnection 类型（指针），并返回
	// *dbConnection 类型。
	// 类型断言 提供了访问接口值底层具体值的方式。
	// t := i.(T)
	// 该语句断言接口值 i 保存了具体类型 T，并将其底层类型为 T 的值赋予变量 t。
	log.Printf("Query: QID[%d] CID[%d]\n", q, conn.(*dbConnection).ID)
}
