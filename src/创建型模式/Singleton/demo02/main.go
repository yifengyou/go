package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creting Single Instance Now")
			singleInstance = &single{}
		} else {
			fmt.Printf("Single Instance already created-1 [%p]\n", singleInstance)
		}
	} else {
		fmt.Printf("Single Instance already created-2 [%p]\n", singleInstance)
	}
	return singleInstance
}


func main() {
	for i := 0; i < 100; i++ {
		go getInstance()
	}
	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()
}