/*
享元模式:
	运用共享技术有效地支持大量细粒度对象。系统只使用少量对象，而这些对象都很相似，状态变化很小，
	可以实现对象地多次复用。由于享元模式要求能够共享地对象必须是细粒度对象，
	因此又称为轻量级模式，是一种结构型模式
关键:

*/

package FlyWeight

type FlyWeight struct {
	Name string
}

func NewFlyWeight(name string) *FlyWeight {
	return &FlyWeight{Name: name}
}

// 享元模式工程类
type FlyWeightFactory struct {
	// 每个string对应一个小对象，实现对象复用
	pool map[string]*FlyWeight
}

// New方法
func NewFlyWeightFactory() *FlyWeightFactory {
	return &FlyWeightFactory{pool: make(map[string]*FlyWeight)}
}

// 获取目标对象
func (f *FlyWeightFactory) GetFlyWeight(name string) *FlyWeight{
	weight,ok := f.pool[name]
	if !ok{
		weight=NewFlyWeight(name)
		f.pool[name]=weight
	}
	return weight
}