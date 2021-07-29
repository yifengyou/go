package Builder

import (
	"fmt"
	"testing"
)

func TestConcreteBuilder_GetResult(t *testing.T) {
	c := HuaWeiComputer{}
	d := Creator{computer: &c}

	// 建造者模式就是支持部分实例化对象
	d.computer.MakeCpu("intel")
	d.computer.MakeKeyBoard("lenovo")
	d.computer.MakeScreen("xiaomi")
	fmt.Printf("%+v\n", c)

	// 当然可以做一个构造器
	computer := d.Construct()
	fmt.Printf("%+v\n", *(computer))
}
