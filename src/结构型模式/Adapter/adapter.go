/*
适配器模式：

关键：

*/
package Adapter

import "fmt"

type Adaptee struct {
}

func (a *Adaptee) SpecificExecute() {
	fmt.Println("充电...")
}

/*
	关键点：适配器嵌套
*/
type Adapter struct {
	*Adaptee
}

func (a *Adapter) Execute() {
	a.SpecificExecute()
}
