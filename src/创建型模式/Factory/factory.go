/*
工厂模式：

*/
package Factory

import "fmt"

// 饭店工厂
type Restaurant interface {
	GetFood()
}

// 饭店工厂之东来顺的模型
type Donglaishun struct {
}

// 模型需要支持工厂必备方法
func (d *Donglaishun) GetFood() {
	fmt.Println("东来顺的饭菜准备继续")
}

// 饭店工厂之庆丰模型
type Qingfeng struct {
}

// 模型需要支持工厂必备方法
func (q *Qingfeng) GetFood() {
	fmt.Println("庆丰包子铺饭菜准备就绪")
}

// 创造不同饭店模板实例
func NewRestaurant(s string) Restaurant {
	switch s {
	case "d":
		return &Donglaishun{}
	case "q":
		return &Qingfeng{}
	}
	return nil
}
