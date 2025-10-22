package main

import (
	"fmt"
	"sync"
)

var totalNum int = 0

var wg sync.WaitGroup

var lock sync.RWMutex

func add() {
	defer wg.Done()
	for range 100000 {
		lock.Lock()
		totalNum++
		lock.Unlock()
	}
	// fmt.Println(totalNum)
}

func sub() {
	defer wg.Done()
	for range 100000 {
		lock.Lock()
		totalNum--
		lock.Unlock()
	}
	// fmt.Println(totalNum)
}

func main() {
	wg.Add(2)
	// 启动协程
	go add()
	go sub()
	wg.Wait()
	fmt.Println(totalNum)
}
