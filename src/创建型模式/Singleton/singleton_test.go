package Sigleton

import (
	"fmt"
	"sync"
	"testing"
)

func TestGetInstance(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(200)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			IncrementAge()
		}()

		go func() {
			defer wg.Done()
			IncrementAge2()
		}()
	}
	wg.Wait()
	p := GetInstance()
	age := p.GetAge()
	fmt.Println(age)
}

func TestGetInstance2(t *testing.T) {
	p1 := GetInstance()
	p1.SetName("p1")
	fmt.Println(p1.GetName()) // p1的Name为"p1"
	p2 := GetInstance()
	p2.SetName("p2")
	fmt.Println(p.GetName())  // p1的Name为"p2"
	fmt.Println(p2.GetName()) // p2的Name为"p2"
	// 接口体指针拷贝(属于不同变量)，结构体指针变量是不同的，但是结构体指针的值相同，指向相同
	fmt.Printf("old:%p  - new:%p\n", &p1, &p2)
	fmt.Printf("old:%p  - new:%p\n", p1, p2)
	// 结构体内部数据指向相同内存区域
	fmt.Printf("old:%p  - new:%p\n", &p1.Name, &p2.Name)
	// 重点，共享实例，本质就是共享结构体，共享内存块
}
