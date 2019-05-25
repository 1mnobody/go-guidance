package main

// 演示去除接口的耦合

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type Data struct {
	Line string
}

type Puller interface {
	Pull(d *Data) error
}

type Storer interface {
	Store(d Data) error
}

type Xenia struct {
}

// Xenia实现了Puller接口
func (Xenia) Pull(d *Data) error {
	switch i := rand.Intn(10); i {
	case 1, 9:
		return io.EOF
	case 5:
		return errors.New("Error Reading Data From Xenia ")
	default:
		d.Line = "Data" + strconv.Itoa(i)
		fmt.Println("In: ", d.Line)
		return nil
	}
}

type Pillar struct {
}

// Pillar 实现了 Storer 接口
func (Pillar) Store(d Data) error {
	fmt.Println("Out:", d.Line)
	return nil
}

type System struct {
	Xenia
	Pillar
}

func pull(p Puller, data []Data) (int, error) {
	for i := range data {
		if err := p.Pull(&data[i]); err != nil {
			return i, err
		}
	}
	return len(data), nil
}

func store(s Storer, data []Data) (int, error) {
	for i, d := range data {
		if err := s.Store(d); err != nil {
			return i, err
		}
	}
	return len(data), nil
}

func Copy(s *System, batch int) error {
	data := make([]Data, batch)
	for {
		i, err := pull(&s.Xenia, data)
		if i > 0{ // 如果pull过程中出错，i就是出错时的那个元素下标，data[:i]切片是一个前闭后开的区间，所以不会包括i
			if _, err := store(&s.Pillar, data[:i]); err != nil {
				return err
			}
		}
		if err != nil {
			return err
		}
	}
}

func main() {
	sys := System{
		Xenia{},
		Pillar{},
	}

	if err := Copy(&sys, 3); err != nil {
		fmt.Println(err)
	}
}