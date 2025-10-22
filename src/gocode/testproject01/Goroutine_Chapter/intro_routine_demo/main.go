package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	for i := range 100 {
		fmt.Println("Hello golang + " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() { // 主线程
	go test() // 开启一个协程
	for i := range 10 {
		fmt.Println("Hello master + " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
