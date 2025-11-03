package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建有缓冲channel，容量为2
	ch := make(chan int, 2)

	// 生产者goroutine
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Printf("发送数据: %d\n", i)
			ch <- i // 发送前3个数据不会阻塞，第4个会阻塞直到有空间
			fmt.Printf("数据 %d 发送完成\n", i)
		}
		close(ch)
	}()

	// 消费者 - 慢速消费
	time.Sleep(3 * time.Second) // 延迟开始消费

	for data := range ch {
		fmt.Printf("接收到数据: %d\n", data)
		time.Sleep(1 * time.Second) // 模拟处理时间
	}
}
