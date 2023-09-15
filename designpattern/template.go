package main

import "fmt"

/*
模板模式

使用场景：针对具有类似操作流程的业务
*/

// IBrew 冲泡饮料接口
type IBrew interface {
	boil()      // 烧水
	cup()       // 取杯
	seasoning() // 放佐料
	pourWater() // 倒水
}

// CoffeeBrew 定义冲咖啡
type CoffeeBrew struct{}

func (cf *CoffeeBrew) boil() {
	fmt.Println("冲咖啡：1、烧水")
}

func (cf *CoffeeBrew) cup() {
	fmt.Println("冲咖啡：2、取杯")
}

func (cf *CoffeeBrew) seasoning() {
	fmt.Println("冲咖啡：3、放入咖啡")
}

func (cf *CoffeeBrew) pourWater() {
	fmt.Println("冲咖啡：4、倒水")
}

func (cf *CoffeeBrew) Brew() {
	fmt.Println("开始冲咖啡啦...")
	cf.boil()
	cf.cup()
	cf.seasoning()
	cf.pourWater()
	fmt.Println("完成冲咖啡啦！！！")
}

// TeaBrew 定义冲茶
type TeaBrew struct{}

func (tea *TeaBrew) boil() {
	fmt.Println("冲茶：1、烧水")
}

func (tea *TeaBrew) cup() {
	fmt.Println("冲茶：2、取杯")
}

func (tea *TeaBrew) seasoning() {
	fmt.Println("冲茶：3、放入茶叶")
}

func (tea *TeaBrew) pourWater() {
	fmt.Println("冲茶：4、倒水")
}

func (tea *TeaBrew) Brew() {
	fmt.Println("开始冲茶啦...")
	tea.boil()
	tea.cup()
	tea.seasoning()
	tea.pourWater()
	fmt.Println("完成冲茶啦！！！")
}

func main() {
	cfBrew := &CoffeeBrew{}
	cfBrew.Brew()
	fmt.Println("*********************")
	teaBrew := &TeaBrew{}
	teaBrew.Brew()
}
