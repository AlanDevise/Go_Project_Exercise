package main

import (
	"fmt"
)

func main() {
	// 只写channel示例
	var intChanWriteOnly chan<- int
	ch := make(chan int, 3) // 创建双向channel
	intChanWriteOnly = ch   // 转换为只写

	intChanWriteOnly <- 20
	intChanWriteOnly <- 30
	fmt.Printf("Channel address: %p\n", intChanWriteOnly)

	// 只读channel示例
	var intChanReadOnly <-chan int
	intChanReadOnly = ch // 同一个channel转换为只读

	// 从只读channel读取
	fmt.Println("Read from read-only channel:", <-intChanReadOnly)
	fmt.Println("Read from read-only channel:", <-intChanReadOnly)

	close(ch) // 关闭channel
}
