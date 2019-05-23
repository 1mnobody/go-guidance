package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

type Runner struct {
	// 接收由操作系统发出的信号
	interrupt chan os.Signal
	// 表示处理完成
	complete chan error
	// 指定耗时
	timeout <-chan time.Time
	// 要执行的函数
	tasks []func(int)
}

var ErrTimeout = errors.New("Received timeout ")
var ErrInterrupt = errors.New("Received interrupt ")

func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

// 变长参数：函数内部接收到的是一个切片
// 传入参数的方式：1、可以将元素一个个列出来 如：Add(task1,task2,task3)；2、传入一个切片，后面加上...，如 tasks... (语法糖，
// 本质其实还是展开切片）
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

func (r *Runner) Start() error {
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		// run会返回error，写到channel中，Start方法会阻塞获取complete\timeout中的数据，直到其中一个
		// channel有数据返回
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete: // run执行完，会写数据到complete这个channel，返回错误信息，
		// 调用Start时要判断err是否为nil（在main的代码中已检查）
		return err
	case <-r.timeout: // timeout中有数据时，说明定时器执行完了
		return ErrTimeout
	}
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.goInterrupt() {
			return ErrInterrupt
		}
		task(id) // tasks 是一个函数列表，这里直接调用函数
	}
	return nil
}

func (r *Runner) goInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}
