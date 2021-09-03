package Factory

import "testing"

func TestNewRestaurant(t *testing.T) {
	// 每个工程实例都是不同模型
	// 每个模型具备工厂定义的方法
	NewRestaurant("d").GetFood()
	NewRestaurant("q").GetFood()
}
