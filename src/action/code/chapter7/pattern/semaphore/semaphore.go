package main

import (
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type (
	// semaphore 是一个接收struct类型的channel，这样定义 semaphore 既是一个channel，也可以实现自定义的方法
	semaphore chan struct{}

	readerWriter struct {
		name           string
		write          sync.WaitGroup
		readerControl  semaphore
		shutdown       chan struct{}
		reportShutdown sync.WaitGroup // 用于等待reader,writer执行完毕
		maxReads       int
		maxReaders     int
		currentReads   int32
	}
)

func init() {
	rand.Seed(time.Now().Unix())
}

// 使用channel实现一个semaphore，允许多个读，只能有一个写
func main() {
	log.Println("Starting Process")
	// 同一时刻最多只能有3个读，有6个reader
	first := start("First", 3, 6)
	second := start("Second", 2, 2)
	time.Sleep(2 * time.Second)
	// shutdown就是调用各个readerWriter的stop方法，stop中会通过rw.reportShutdown 来等待所有的reader以及Writer执行完毕
	shutdown(first, second)
	log.Println("Process End")
	return
}

func start(name string, maxReads int, maxReaders int) *readerWriter {
	rw := readerWriter{
		name:       name,
		shutdown:   make(chan struct{}),
		maxReads:   maxReads,
		maxReaders: maxReaders,
		// readerControl 控制reader的数量（最多maxReads个）
		readerControl: make(semaphore, maxReads),
	}

	// 在shutdown中，reportShutdown.Wait()，等待Reader和Writer执行完毕
	rw.reportShutdown.Add(maxReaders)
	// readerWriter 创建 maxReaders 个 goroutine
	for goroutine := 0; goroutine < maxReaders; goroutine++ {
		go rw.reader(goroutine)
	}

	rw.reportShutdown.Add(1)
	go rw.writer()

	return &rw
}

func shutdown(writer ...*readerWriter) {
	var wg sync.WaitGroup
	wg.Add(len(writer))
	for _, readerWriter := range writer {
		go readerWriter.stop(&wg)
	}

	wg.Wait()
}

func (rw *readerWriter) stop(group *sync.WaitGroup) {
	defer group.Done()

	log.Printf("%s\t: #####> Stop", rw.name)

	// close掉的channel，在select中也是可执行的
	close(rw.shutdown)

	// 等待readerWriter的所有goroutine执行完成
	rw.reportShutdown.Wait()

	log.Printf("%s\t: #####> Stopped", rw.name)
}

func (rw *readerWriter) reader(reader int) {
	defer rw.reportShutdown.Done()

	for {
		select {
		case <-rw.shutdown: // rw.shutdown关闭时，这个case会被执行
			log.Printf("%s\t: #> Reader Shutdown", rw.name)
			return
		default:
			rw.performRead(reader)
		}
	}
}

func (rw *readerWriter) performRead(reader int) {
	rw.ReadLock(reader)
	count := atomic.AddInt32(&rw.currentReads, 1)
	log.Printf("%s\t: [%d] Start\t- [%d] Reads\n", rw.name, reader, count)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	count = atomic.AddInt32(&rw.currentReads, -1)
	log.Printf("%s\t: [%d] Finish\t- [%d] Reads\n", rw.name, reader, count)

	rw.ReadUnlock(reader)
}

func (rw *readerWriter) writer() {
	defer rw.reportShutdown.Done()
	for {
		select {
		case <-rw.shutdown: // rw.shutdown关闭时，这个case会被执行
			log.Printf("%s\t: #> Writer Shutdown", rw.name)
			return
		default:
			rw.performWrite()
		}
	}
}

func (rw *readerWriter) performWrite() {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("%s\t: *****> Writing Pending\n", rw.name)
	rw.WriteLock()

	log.Printf("%s\t: *****> Writing Start", rw.name)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("%s\t: *****> Writing Finish", rw.name)

	rw.WriteUnlock()
}

func (rw *readerWriter) ReadLock(reader int) {
	// write 是一个waitGroup，这里等待就是等待写完成（写之前会先进入rw.write.Add(1)，写完成后会Done）
	rw.write.Wait()
	// Acquire 方法往 semaphore （chan struct{}）中写数据，缓冲区写满之后，再写入会被阻塞，
	// 故缓冲区大小就是能并发读（readerWriter.maxReads）的数量
	rw.readerControl.Acquire(1)
}

func (rw *readerWriter) ReadUnlock(reader int) {
	// Release 就是从 semaphore中读数据，释放掉semaphore中的缓冲区
	rw.readerControl.Release(1)
}

func (rw *readerWriter) WriteLock() {
	// 这里Add(1)，在ReadLock()中会Wait，从而进入写之后，新的read操作无法再进行
	rw.write.Add(1)
	// 写满缓冲区，如果存在正在读的routine,这里也会被阻塞，直到写满缓冲区，写满后，再次调用Acquire()也会被阻塞
	rw.readerControl.Acquire(rw.maxReads)
}

func (rw *readerWriter) WriteUnlock() {
	rw.readerControl.Release(rw.maxReads)
	rw.write.Done()
}

func (s semaphore) Acquire(buffers int) {
	var e struct{}
	for buffer := 0; buffer < buffers; buffer++ {
		s <- e
	}
}

func (s semaphore) Release(buffers int) {
	for buffer := 0; buffer < buffers; buffer++ {
		<-s
	}
}
