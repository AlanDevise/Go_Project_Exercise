package main

import "fmt"

// 主函数入口
func main() {
	result := cal(10, 20)
	fmt.Println(result)

	result = cal(30, 50)
	fmt.Println(result)

	// 利用下划线忽略某一个返回值
	_, diff := cal_2(50, 20)
	fmt.Println(diff)
}

// 计算两数之和
// func 函数名 （形参）（返回参数）
// 首字母大写的函数可以被其他包调用，首字母小写只能在本包内调用
func cal(num1 int, num2 int) (sum int) {
	sum = num1 + num2
	return sum
}

// 计算两数之和与两数之差
func cal_2(num1 int, num2 int) (sum int, diff int) {
	sum = num1 + num2
	diff = num1 - num2
	return sum, diff
}
