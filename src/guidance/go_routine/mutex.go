package go_routine

import (
	"fmt"
	"sync"
	"time"
)

func MutexDemo() {
	m := &sync.Mutex{}
	var i int32 = 10
	go Inc(m, &i)
	// 休眠5ms，保证Inc先执行
	time.Sleep(5 * time.Millisecond)
	go Value(m, &i)
}

// 互斥锁 sync.Mutex，两个方法Lock, Unlock。可借助defer语句来保证Unlock一定会执行
func Inc(mutex *sync.Mutex, num *int32) {
	fmt.Printf("%p start inc ...\n", mutex)
	mutex.Lock()
	time.Sleep(3 * time.Second)
	*num++
	mutex.Unlock()
}

func Value(mutex *sync.Mutex, num *int32) int32 {
	mutex.Lock()
	defer mutex.Unlock()
	fmt.Printf("%p %v\n", mutex, *num)
	return *num
}
