/*
建造者模式：
	将一个复杂对象的构建与它的表示分离，使得同样的构建过程可以创建不同的表示
	Builder模式适用于这么一种情况：无法或不想一次性把实例的所有属性都给出，
	而是要分批次、分条件构造
关键：
	分批次、分条件构造 - 细粒度构造
*/
package Builder

import "fmt"

//电脑对象的建造
type Computer interface {
	//主板制造
	MakeCpu(string)
	//键盘制造
	MakeKeyBoard(string)
	//屏幕制造
	MakeScreen(string)
}

// 定义一个构建器
type Creator struct {
	computer Computer
}

// 创建一个电脑
func (d *Creator) Construct() *Computer {
	d.computer.MakeCpu("origin cpu")
	d.computer.MakeKeyBoard("origin keyboard")
	d.computer.MakeScreen("origin screen")
	return &d.computer
}

// 构建一个具体品牌的电脑
type HuaWeiComputer struct {
	Cpu      string
	KeyBoard string
	Screen   string
}

func (l *HuaWeiComputer) MakeCpu(cpuName string) {
	fmt.Println("主板构建中...")
	l.Cpu = cpuName
}
func (l *HuaWeiComputer) MakeKeyBoard(keyboardName string) {
	fmt.Println("键盘构建中...")
	l.KeyBoard = keyboardName
}
func (l *HuaWeiComputer) MakeScreen(screenName string) {
	fmt.Println("屏幕构建中...")
	l.Screen = screenName
}
