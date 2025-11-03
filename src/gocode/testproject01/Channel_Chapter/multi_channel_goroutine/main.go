package main

import (
	"fmt"
	"sync"
	"time"
)

// 只需定义无需赋值
var waitGroup sync.WaitGroup

// 写
func writeDate(intChan chan int) {
	defer waitGroup.Done()
	for i := 1; i < 10; i++ {
		intChan <- i
		fmt.Println("写入的数据为：", i)
		// 写完数据停一下
		time.Sleep(time.Second)
	}
	// 管道关闭
	close(intChan)
}

// 读
func readData(intChan chan int) {
	defer waitGroup.Done()
	// 遍历
	for v := range intChan {
		fmt.Println("读取的数据为：", v)
		time.Sleep(time.Second)
	}
}

func main() {
	// 写协程和读协程共同操作的一个管道，定义这个管道
	intChan := make(chan int, 10)

	// 开启读写协程
	waitGroup.Add(2)
	go writeDate(intChan)
	go readData(intChan)

	// 主线程阻塞直到两个协程执行完毕
	waitGroup.Wait()
}
