/*
迭代器模式（Iterator Design Pattern）:
	游标模式（Cursor Design Pattern）
	迭代器模式将集合对象的遍历操作从集合类中拆分出来，放到迭代器类中，让两者的职责更加单一

关键：


参考：


*/
package Iterator

type Iterator interface {
	Index() int
	Value() interface{}
	HasNext() bool
	Next()
}

type ArrayIterator struct {
	array []interface{}
	index *int
}

// 获取索引
func (a *ArrayIterator) Index() *int {
	return a.index
}

// 获取值，具体是哪个索引？迭代器索引与数据放在一起
func (a *ArrayIterator) Value() interface{} {
	return a.array[*a.index]
}

// 判断是否还有下一个元素
func (a *ArrayIterator) HasNext() bool {
	return *a.index+1 <= len(a.array)
}

// 更新索引，切换到下一个元素
func (a *ArrayIterator) Next() {
	if a.HasNext() {
		*a.index++
	}
}
