/*
外观（门面）设计模式:
	内部子系统的任何变化不会影响到Facade接口的变化。
	Facade 设计模式更注重从架构的层次去看整个系统，而不是单个类的层次。Facade 很多时候更是一种架构设计模式。

关键：

*/
package Facade

import "fmt"

type CarModel struct {
}

func NewCarModel() *CarModel {
	return &CarModel{}
}

func (c *CarModel) SetModel() {
	fmt.Printf(" CarModel - SetModel\n")
}

type CarEngine struct {
}

func NewCarEngine() *CarEngine {
	return &CarEngine{}
}

func (c *CarEngine) SetEngine() {
	fmt.Printf(" CarEngine - SetEngine\n")
}

type CarBody struct {
}

func NewCarBody() *CarBody {
	return &CarBody{}
}

func (c *CarBody) SetBody() {
	fmt.Printf(" CarBody - SetBody\n")
}

type CarAccessories struct {
}

func NewCarAccessories() *CarAccessories {
	return &CarAccessories{}
}

func (c *CarAccessories) SetAccessories() {
	fmt.Printf(" CarAccessories - SetAccessories\n")
}

// 关键：定义门面结构体
type CarFacade struct {
	accessories *CarAccessories
	body        *CarBody
	engine      *CarEngine
	model       *CarModel
}

// 门面实例化方法
func NewCarFacade() *CarFacade {
	return &CarFacade{
		accessories: NewCarAccessories(),
		body:        NewCarBody(),
		engine:      NewCarEngine(),
		model:       NewCarModel(),
	}
}

// 使用门面创建完整后端，对应接口
func (c *CarFacade) CreateCompleteCar() {
	fmt.Printf("******** Creating a Car **********\n")
	c.model.SetModel()
	c.engine.SetEngine()
	c.body.SetBody()
	c.accessories.SetAccessories()
	fmt.Printf("******** Car creation is completed. **********\n")
}
