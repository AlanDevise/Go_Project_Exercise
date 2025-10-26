package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 加入读写锁
var lock sync.RWMutex

func read() {
	defer wg.Done()
	// 加读锁
	lock.RLock()
	fmt.Println("开始读取数据")
	time.Sleep(time.Second)
	fmt.Println("读取数据成功")
	// 解读锁
	lock.RUnlock()
}

func write() {
	defer wg.Done()
	// 加写锁
	lock.Lock()
	fmt.Println("开始写入数据")
	for i := 0; i < 6; i++ {
		fmt.Println("写入中...")
		time.Sleep(time.Second)
	}
	fmt.Println("写入数据成功")
	// 解写锁
	lock.Unlock()
}

func main() {
	// 这里一定要注意Add方法的入参需要和实际的协程数量一致，否则可能会出现panic
	wg.Add(6)
	// 启动协程----场景：读多写少
	for i := 0; i < 5; i++ {
		go read()
	}
	go write()
	wg.Wait()

}
