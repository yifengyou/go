package Prototype

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConcretePrototype_Clone(t *testing.T) {
	name := "出去浪"
	proto := ConcretePrototype{StructName: name}
	newProto := proto.Clone()
	actualName := newProto.Name()
	// 使用原型模式拷贝的，创建了新对象，变量地址不同，结构体内部数据地址不同
	fmt.Printf("old:%p  - new:%p\n", &proto, &newProto)
	// 属于全新的结构体，仅仅是基于一个模板clone出一个新的结构体
	// 只是这个clone需要用户提供
	fmt.Printf("old:%p  - new:%p\n", &proto.StructData,
		&newProto.(*ConcretePrototype).StructData)
	assert.Equal(t, name, actualName)
}
