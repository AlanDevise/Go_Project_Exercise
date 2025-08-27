package main

import "fmt"

func main() {
	test()
	fmt.Println("上面的触发执行成功")
	fmt.Println("正常执行下面的逻辑")
}

func test() {
	// 利用defer+recover来捕获错误，防止程序崩溃，defer后面加上匿名函数
	defer func() {
		err := recover()
		// 如果没有捕获错误，返回值为零值：nil
		if err != nil {
			fmt.Println("捕获到错误:", err)
		}
	}()
	num1 := 10
	num2 := 0
	result := num1 / num2
	fmt.Println("结果是:", result)
}
