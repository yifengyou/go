package Iterator

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	array := []interface{}{1, 3, 9, 2, 8}
	a := 0
	iterator := ArrayIterator{array: array, index: &a}
	// 迭代器内置索引、游标
	// 是否有下一个元素
	// 切换到下一个元素
	for it := iterator; iterator.HasNext(); iterator.Next() {
		// 获取当前索引，获取当前值
		index, value := it.Index(), it.Value().(int)
		if value != array[*index] {
			fmt.Println("error...")
		}
		fmt.Println(*index, value)
	}
}
