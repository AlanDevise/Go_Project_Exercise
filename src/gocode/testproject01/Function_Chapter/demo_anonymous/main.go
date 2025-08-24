package main

import "fmt"

// 主函数入口
func main() {
	// 定义一个匿名函数并立即调用
	result := func(num1 int, num2 int) int {
		return num1 + num2
	}(10, 20)
	fmt.Println("The result is:", result)
}
