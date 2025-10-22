package main

import (
	"fmt"
	"sync"
)

// 只用定义不需要赋值
var wg sync.WaitGroup

func main() {
	var count = 5
	wg.Add(count) // 知道次数的话，可以提前一次性Add
	// 启动5个协程
	for i := 1; i <= count; i++ {
		// wg.Add(1) // 协程开始的时候+1
		go func(n int) {
			// 防止忘记计数器减1，可以结合defer关键字
			defer wg.Done() // 协程执行完成减1
			fmt.Println(n)
		}(i)
	}
	// 主线程一直在这里阻塞，直到wg减为0，就停止
	wg.Wait()
}
