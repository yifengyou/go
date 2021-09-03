package Abstract_Factory

import "testing"

func TestNewSimpleShapeFactory(t *testing.T) {
	// 创造抽象工程
	factory := NewSimpleShapeFactory()

	// 创造food
	food := factory.CreateFood()
	food.Cook()

	// 创造vegetable
	vegetable := factory.CreateVegetable()
	vegetable.Cook()
}
