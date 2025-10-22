package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		// 匿名函数，直接调用匿名函数
		go func(n int) {
			fmt.Println(n)
		}(i)
	}
	time.Sleep(time.Second * 5)
}
