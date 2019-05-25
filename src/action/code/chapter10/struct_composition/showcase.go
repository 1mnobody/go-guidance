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

type Xenia struct {

}

func (Xenia) Pull(d *Data) error {
	switch i := rand.Intn(10); i {
	case 1,9 :
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

// System 组合了两个struct, 从Xenia中 Pull 数据，往Pillar中Store数据
type System struct {
	Xenia
	Pillar
}

func pull(x *Xenia, data []Data) (int, error) {
	for i := range data {
		if err := x.Pull(&data[i]); err != nil {
			return i, err
		}
	}
	return len(data), nil
}

func store(p *Pillar, data []Data) (int, error) {
	for i, d := range data {
		if err := p.Store(d); err != nil {
			return i, err
		}
	}
	return len(data), nil
}

func Copy(sys *System, batch int) error {
	data := make([]Data, batch)
	for {
		i, err := pull(&sys.Xenia, data)
		if i > 0 {// 如果pull过程中出错，i就是出错时的那个元素下标，data[:i]切片是一个前闭后开的区间，所以不会包括i
			if _, err := store(&sys.Pillar, data[:i]); err != nil {
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