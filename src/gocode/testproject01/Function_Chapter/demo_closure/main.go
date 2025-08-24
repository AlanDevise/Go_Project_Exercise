package main

import "fmt"

// 函数功能：求和
// 函数的名字：getSum 参数为空
// getSum这个函数返回值是一个函数，这个函数有一个int类型的参数，返回值是int类型
func getSum() func(int) int {
	var sum int = 0
	return func(num int) int {
		sum = sum + num
		return sum
	}
}

// 返回的这个匿名函数引用了外部的sum变量，形成了闭包
// 闭包：函数+引用环境变量

// 主函数入口
func main() {
	oneFunction := getSum()
	fmt.Println(oneFunction(1)) // 10
	fmt.Println(oneFunction(2)) // 10
	fmt.Println(oneFunction(3)) // 10
}
