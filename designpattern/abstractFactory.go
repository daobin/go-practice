package main

/**
抽象工厂方法模式
*/

// 抽象层

type AbstractGpu interface {
	ShowGpu()
}

type AbstractCpu interface {
	ShowCpu()
}

type AbstractMemory interface {
	ShowMemory()
}

type AbstractFactory interface {
	BuyGpu() AbstractGpu
	BuyCpu() AbstractCpu
	BuyMemory() AbstractMemory
}

// 实现层

type IntelGpu struct{}

func (i *IntelGpu) ShowGpu() {
	println("Intel Gpu")
}

type IntelCpu struct{}

func (i *IntelCpu) ShowCpu() {
	println("Intel Cpu")
}

type IntelMemory struct{}

func (i *IntelMemory) ShowMemory() {
	println("Intel Memory")
}

type IntelFactory struct{}

func (i *IntelFactory) BuyGpu() AbstractGpu {
	return new(IntelGpu)
}
func (i *IntelFactory) BuyCpu() AbstractCpu {
	return new(IntelCpu)
}
func (i *IntelFactory) BuyMemory() AbstractMemory {
	return new(IntelMemory)
}

// 业务处理
func main() {
	// 组装电脑：Intel的CPU、Intel的显卡、Intel的内存
	intelFactory := new(IntelFactory)
	intelFactory.BuyCpu().ShowCpu()
	intelFactory.BuyGpu().ShowGpu()
	intelFactory.BuyMemory().ShowMemory()

	// 组装电脑：Intel的CPU、Nvidia的显卡、Kingston的内存
}
