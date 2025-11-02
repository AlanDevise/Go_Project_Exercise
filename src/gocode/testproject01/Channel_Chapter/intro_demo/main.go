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

	fmt.Printf("管道的实际长度：%v，管道的容量是：%v\n", len(intChan), cap(intChan))

	// 从管道中读取数据
	num1 := <-intChan
	fmt.Printf("当前num1的值为：%v\n", num1)
	// 输出管道长度：
	fmt.Printf("管道的实际长度：%v，管道的容量是：%v\n", len(intChan), cap(intChan))

	// 从管道中读取数据
	num2 := <-intChan
	fmt.Printf("当前num1的值为：%v\n", num2)
	// 输出管道长度：
	fmt.Printf("管道的实际长度：%v，管道的容量是：%v\n", len(intChan), cap(intChan))

	// 从管道中读取数据
	num3 := <-intChan
	fmt.Printf("当前num1的值为：%v\n", num3)
	// 输出管道长度：
	fmt.Printf("管道的实际长度：%v，管道的容量是：%v\n", len(intChan), cap(intChan))

	// 从管道中读取数据
	num4 := <-intChan
	fmt.Printf("当前num1的值为：%v\n", num4)
	// 输出管道长度：
	fmt.Printf("管道的实际长度：%v，管道的容量是：%v\n", len(intChan), cap(intChan))
}
