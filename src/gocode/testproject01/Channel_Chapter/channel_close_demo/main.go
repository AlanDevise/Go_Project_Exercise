package main

import (
	"fmt"
)

func main() {
	// 声明一个管道 定义一个int类型的管道
	var intChan chan int
	// 通过make进行初始化，管道可以存放3个int类型的数据
	intChan = make(chan int, 3)
	// 证明逛到是引用类型：形如 0x14000020090
	fmt.Printf("intChan的值: %v\n", intChan)

	// 向管道存放数据：
	intChan <- 10
	num := 20
	intChan <- num
	intChan <- 40

	// 关闭管道
	close(intChan)

	// 再次写入数据:-->  出错，已关闭管道不可继续写入
	// intChan <- 888

	// 当管道关闭后，读取数据是可以的
	num1 := <-intChan
	fmt.Println(num1)
}
