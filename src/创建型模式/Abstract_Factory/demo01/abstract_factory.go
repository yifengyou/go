/*
抽象工厂模式：
	提供一个创建一系列相关或相互依赖对象的接口，而无须指定它们具体的类
关键：
	抽象接口

 */
package Abstract_Factory

import "fmt"

type Lunch interface {
	Cook()
}

// 工厂产品rise
type rise struct {
}

func (c *rise) Cook() {
	fmt.Println("cook rise.")
}

// 工厂产品tomato
type Tomato struct {
}

func (c *Tomato) Cook() {
	fmt.Println("cook tomato.")
}

// 抽象工程接口，关键
type LunchFactory interface {
	CreateFood() Lunch
	CreateVegetable() Lunch
}

// 抽象工程结构体
type simpleLunchFactory struct {
}

// 工厂实例化函数
func NewSimpleShapeFactory() LunchFactory {
	return &simpleLunchFactory{}
}

// 工程方法
func (s *simpleLunchFactory) CreateFood() Lunch {
	return &rise{}
}

// 工程方法
func (s *simpleLunchFactory) CreateVegetable() Lunch {
	return &Tomato{}
}
