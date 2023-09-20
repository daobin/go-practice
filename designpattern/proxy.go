package main

import "fmt"

/**
代理模式

要增加“自己”的功能，不想做出改变，就选择一位“代理”来完成，符合开闭原则思想
*/

type Goods struct {
	Kind string // 商品种类
	Fact bool   // 商品真伪
}

// 抽象层

type Shopping interface {
	Buy(goods *Goods)
}

// 实现层

type KoreaShopping struct{}

func (s *KoreaShopping) Buy(goods *Goods) {
	fmt.Println("韩国购物：", goods.Kind)
}

type AmericanShopping struct{}

func (s *AmericanShopping) Buy(goods *Goods) {
	fmt.Println("美国购物：", goods.Kind)
}

// 购物代理

type ShoppingProxy struct {
	shopping Shopping
}

func (s *ShoppingProxy) Buy(goods *Goods) {
	if s.CheckFact(goods) {
		s.shopping.Buy(goods)
	}
}

func (s *ShoppingProxy) CheckFact(goods *Goods) bool {
	fmt.Println("对商品【", goods.Kind, "】辨别真伪：", goods.Fact)
	return goods.Fact
}

func NewProxy(shopping Shopping) Shopping {
	return &ShoppingProxy{shopping: shopping}
}

func main() {
	gd1 := &Goods{
		Kind: "面膜",
		Fact: true,
	}
	gd2 := &Goods{
		Kind: "CET4证书",
		Fact: false,
	}

	shopping := new(KoreaShopping)
	proxy := NewProxy(shopping)
	proxy.Buy(gd1)
	proxy.Buy(gd2)
}
