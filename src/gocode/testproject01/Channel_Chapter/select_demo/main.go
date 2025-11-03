package main

import (
	"fmt"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

func main() {

	waitGroup.Add(2)

	// 定义一个int管道：
	intChan := make(chan int, 1)
	go func() {
		defer waitGroup.Done()
		time.Sleep(time.Second * 5)
		intChan <- 10
	}()

	// 定一个一个string管道
	stringChan := make(chan string, 1)
	go func() {
		defer waitGroup.Done()
		time.Sleep(time.Second * 2)
		stringChan <- "helloworld"
	}()

	// fmt.Println("value from intChan", <-intChan)

	select {
	case value := <-intChan:
		fmt.Println("intChan:", value)
	case value := <-stringChan:
		fmt.Println("stringChan:", value)
	default:
		fmt.Println("兜底逻辑")
	}
	waitGroup.Wait()
}
