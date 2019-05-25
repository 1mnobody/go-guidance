package main

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

// 相当于Java中 PullStorer 是 Puller 与 Storer两个接口的子接口
type PullStorer interface {
	Puller
	Storer
}

type Xenia struct {
}

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

func (Pillar) Store(d Data) error {
	fmt.Println("Out:", d.Line)
	return nil
}

// System 组合了两个struct ，实现了上面的三个接口：Puller, Storer, PullStorer
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

// ps 是 PullStorer 接口，而 PullStorer 既是Puller接口，也是Storer接口
func  Copy(ps PullStorer, batch int) error {
	data := make([]Data, batch)
	for {
		i, err := pull(ps, data)
		if i > 0 {
			if _, err := store(ps, data[:i]); err != nil {
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
