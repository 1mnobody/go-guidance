package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

type Pool struct {
	m         sync.Mutex
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed    bool
}

var ErrPoolClosed = errors.New("Pool has been closed. ")

// 用于创建 Pool，传入的参数中，fn（一个函数）用于创建新的资源，size则是指定pool的大小（这里的实现其实
// 就是指定了channel的缓冲区大小）
func New(fn func() (c io.Closer, err error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("Size value too small. ")
	}
	return &Pool{
		resources: make(chan io.Closer, size),
		factory:   fn,
	}, nil
}

func (p *Pool) Acquire() (io.Closer, error) {
	select {
	// 首先尝试从channel的缓冲区中拿一个Closer实例（一个资源），如果channel中无数据，则执行default
	case r, ok := <-p.resources:
		log.Println("Acquire:", "Shared Resource")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default: // 调用factory() 新建一个资源
		log.Println("Acquire", "New Resource")
		return p.factory()
	}
}

func (p *Pool) Release(r io.Closer) {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		r.Close()
		return
	}

	select {
	case p.resources <- r: // 资源池没满，将资源放回资源池（channel的缓冲区）
		log.Println("Release Resource to Queue(channel)", r)
	default: // 资源池满，直接关闭多余的资源数据
		log.Println("Release, Closing", r)
		r.Close()
	}
}

func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		return
	}

	p.closed = true

	close(p.resources)
	// channel close之后，依旧可以从channel中获取数据，但是不能再发送数据到channel中
	for r := range p.resources {
		r.Close()
	}
}
