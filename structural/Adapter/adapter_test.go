package Adapter

import "testing"

func TestAdaptee_SpecificExecute(t *testing.T) {
	adapter := Adapter{}
	// adapter 调用 adaptee的Execute方法
	adapter.Execute()
}