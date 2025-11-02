package main

import (
	"fmt"
)

func main() {
	// 声明一个管道 定义一个int类型的管道
	var intChan chan int
	// 通过make进行初始化，管道可以存放3个int类型的数据
	intChan = make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan <- i
	}

	// 写入完毕后关闭通道，告知读取方“不会再有新数据”
	close(intChan)

	// 开始遍历前如果不关闭channel，会出现deadlock的现象
	// 遍历 for-range
	for v := range intChan {
		fmt.Println("value= ", v)
	}

}
