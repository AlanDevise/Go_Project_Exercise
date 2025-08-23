package main

import (
	"fmt"
	"strconv"
)

// 字符串示例
func main() {
	var c1 byte = 'c'
	fmt.Println("c1=", c1)

	var c2 byte = ')'
	fmt.Println("c2=", c2)

	// 字符串和字符是两种不一样的数据类型，下面的代码就是错误的
	// var c3 byte = "s"
	// fmt.Println("c3=", c3)

	var c4 = '中'
	fmt.Println("c4=", c4)
	fmt.Printf("c4=%T\n", c4)

	var str1 string = "你好我叫Alan"
	fmt.Println("str1=", str1)
	fmt.Printf("str1=%T\n", str1)

	// 多行字符串
	// 反引号``
	// 可以实现字符串的多行输入，且不需要转义
	// 反引号中间的内容原样输出
	// 反引号中间不能出现反引号
	var str2 string = `
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
	`
	fmt.Println("str2=", str2)
	fmt.Printf("str2=%T\n", str2)

	// string与其他类型之间的转换，有且仅有强制类型转换
	var var3 string = "10.9"
	var str4 = string(var3)
	fmt.Println("str4=", str4)
	fmt.Printf("str4=%T\n", str4)


	var str5 string = "true"
	var isOk bool = false
	isOk, _ = strconv.ParseBool(str5)
	fmt.Println("isOk=", isOk)
	fmt.Printf("isOk=%T\n", isOk)
}