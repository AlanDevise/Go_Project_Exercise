package main

import "fmt"

var globalVar = "全局变量"



func main() {
	var name = 10
	fmt.Println("main函数开始执行", name)
	var xiaoshu = 12.34
	fmt.Println("xiaoshu=", xiaoshu)
	fmt.Printf("xiaoshu的类型是%T\n", xiaoshu)
	var isOk = true
	fmt.Println("isOk=", isOk)
	var b = 10
	fmt.Println("b=", b)
	fmt.Printf("b的类型是%T\n", b)

	var variable1 int
	fmt.Println("variable1=", variable1)

	// 一次性多个变量声明
	var (
		variable2  = 500
		variable3  = "hello"
	)

	fmt.Println("globalVar=", globalVar)
	fmt.Println("variable2=", variable2)
	fmt.Println("variable3=", variable3)
}