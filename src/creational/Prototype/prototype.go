/*
原型模式：
	原型模式使用原型实例指定创建对象的种类,并且通过拷贝原型对象创建新的对象
	原型模式是一种比较简单的模式，也非常容易理解，
	实现一个接口，重写一个方法即完成了原型模式。
	原型模式并不直接通过类或者结构体来实例化，而是通过一个实例对自身进行clone来得到一个新的实例
	(其实一般情况也就是 clone 方法自己偷偷的实例化了一个对象然后把属性copy过去)，
	原型模式和直接实例化的最大区别就是通过原型模式，可以直接把实例clone时 自身的状态也一起copy过去
关键：
	状态拷贝
*/
package Prototype

type Prototype interface {
	Name() string
	Clone() Prototype
}

type ConcretePrototype struct {
	StructName string
	StructData int
}

func (p *ConcretePrototype) Name() string {
	return p.StructName
}

func (p *ConcretePrototype) Clone() Prototype {
	return &ConcretePrototype{StructName: p.StructName}
}
