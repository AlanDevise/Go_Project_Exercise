package main

import (
	"fmt"
	"sync"
)

var waitGroup sync.WaitGroup

// 输出数字
func printNum() {
	defer waitGroup.Done()
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

// 做除法操作
func device() {
	defer waitGroup.Done()
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("devide()异常：", err)
		}
	}()
	num1 := 10
	num2 := 0
	result := num1 / num2
	fmt.Println(result)
}

func main() {
	waitGroup.Add(2)
	go printNum()
	go device()
	waitGroup.Wait()
}
