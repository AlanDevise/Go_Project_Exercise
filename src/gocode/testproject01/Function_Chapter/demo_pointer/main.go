package main

import "fmt"

// 主函数入口
func main() {
	var a int = 10
	var b int = 20
	var c int = 30
	var str1 = "hello"
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("c =", c)
	// &符号获取指定变量的内存地址
	fmt.Println("&a =", &a)

	// 定义一个指针变量，var表示声明一个变量，ptr 指针变量的名字，ptr对应的类型是：*int 是一个指针类型，是一个指向int类型的指针
	var ptr *int = &a
	fmt.Println("ptr =", ptr)   // 取指针变量的值，指针变量的值是一个内存地址
	fmt.Println("*ptr =", *ptr) // 取指针变量指向的内存地址的值，*ptr = 10
	fmt.Println("&ptr =", &ptr) // 指针变量自己的内存地址
	// fmt.Println("ptr指向的数值 =", *&*&*&*&*&*&*&*&*&*&*&*&*&*&*&*&*&*&*&*&*&*&*&*&*&*&*&*&*&*&*&*&*&ptr) // 指针变量自己的内存地址

	var ptr2 *string = &str1
	fmt.Println("ptr2 =", ptr2)   // 取指针变量的值，指针变量的值是一个内存地址
	// ptr = ptr2 //  编译报错：cannot use ptr2 (type *string) as type *int in assignment

	// if的括号可以去掉
	if a < 20 {
		fmt.Println("a的值小于20")
	}

	if b := 45; b < 20 {
		fmt.Println("b的值小于20")
	} else {
		fmt.Println("b的值大于等于20")
	}
}