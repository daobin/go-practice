package main

import "fmt"

/**
装饰模式
*/

// 抽象层

type Phone interface {
	Show()
}

type Decorator struct {
	phone Phone
}

// 实现层

type Huawei struct{}

func (h *Huawei) Show() {
	fmt.Println("秀手机：华为")
}

type Xiaomi struct{}

func (x *Xiaomi) Show() {
	fmt.Println("秀手机：小米")
}

type MoDecorator struct {
	Decorator
}

func (m *MoDecorator) Show() {
	m.phone.Show()
	fmt.Println("贴膜的手机")
}

func NewMoDecorator(phone Phone) Phone {
	return &MoDecorator{Decorator{phone}}
}

type KeDecorator struct {
	Decorator
}

func (k *KeDecorator) Show() {
	k.phone.Show()
	fmt.Println("有壳的手机")
}

func NewKeDecorator(phone Phone) Phone {
	return &KeDecorator{Decorator{phone}}
}

func main() {
	hw := new(Huawei)
	hw.Show()
	fmt.Println("----------------")

	moHw := NewMoDecorator(hw)
	moHw.Show()
	fmt.Println("----------------")

	keHw := NewKeDecorator(hw)
	keHw.Show()
}
