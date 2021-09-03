/*
代理模式：

关键：

 */
package Proxy

import (
	"fmt"
	"strconv"
)

//type ITask interface {
//	RentHouse(desc string, price int)
//}

// 被代理对象及其方法
type Task struct {
}

func (t *Task) RentHouse(desc string, price int) {
	fmt.Sprintln(fmt.Printf("租房子的地址%s,价钱%s,中介费%s.", desc, strconv.Itoa(price), strconv.Itoa(price)))
}

// 代理对象及其方法
type AgentTask struct {
	task *Task
}

func NewAgentTask() *AgentTask {
	return &AgentTask{task: &Task{}}
}

func (t *AgentTask) RentHouse(desc string, price int) {
	t.task.RentHouse(desc, price)
}
