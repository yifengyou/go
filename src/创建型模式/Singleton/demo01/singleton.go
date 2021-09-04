/*
单例模式
	单例模式可以保证系统中一个类只有一个实例且该实例易于外界访问，
	从而方便对实例个数的控制并节约系统资源。
	为了做到线程安全，对单例对象的修改需要加锁
*/
package Sigleton

import "sync"

// 包中定义全局变量
// sync.Once 使用变量 done 来记录函数的执行状态，
// 使用 sync.Mutex 和 sync.atomic 来保证线程安全的读取 done
// 参数为func，即使更换func也只会执行最开始的func
var (
	p    *Pet
	once sync.Once
)

// 可以放到init中，默认导入包就实例化，也可以单独开来
func init() {
	once.Do(func() {
		p = &Pet{}
	})
}

// 如果是New，需要返回对象
func New() *Pet {
	once.Do(func() {
		p = &Pet{}
	})
	return p
}

func GetInstance() *Pet {
	return p
}

type Pet struct {
	Name string
	age  int
	mux  sync.Mutex
}

func (p *Pet) SetName(n string) {
	p.mux.Lock()
	defer p.mux.Unlock()
	p.Name = n
}

func (p *Pet) GetName() string {
	p.mux.Lock()
	defer p.mux.Unlock()
	return p.Name
}

func (p *Pet) IncrementAge() {
	p.mux.Lock()
	p.age++
	p.mux.Unlock()
}

func (p *Pet) GetAge() int {
	p.mux.Lock()
	defer p.mux.Unlock()
	return p.age
}
