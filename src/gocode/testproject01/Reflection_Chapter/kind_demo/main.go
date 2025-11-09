package main

import (
	"fmt"
	"reflect"
)

// 利用一个函数，函数的参数定义为空接口
func testReflect(i interface{}) {
	// 空接口没有任何方法，所以可以理解为所有类型都实现了空接口，也可以理解为我们可以把任何一个变量赋给空接口
	// 1. 调用typeOf函数，返回reflect.Type 类型
	reType := reflect.TypeOf(i)
	fmt.Println("retype:", reType)
	// 2. 调用valueOf函数，返回reflect.Vlaue类型
	reValue := reflect.ValueOf(i)
	fmt.Println("reValue:", reValue)

	// 获取变量的类别
	kind1 := reType.Kind()
	fmt.Println(kind1)

	// 获取变量类型
	if n, isStudent := reValue.Interface().(Student); isStudent {
		fmt.Printf("结构体的类型是：%T\n", n)
	}

}

type Student struct {
	Name string
	Age  int
}

func main() {
	// 对结构体类型进行反射
	// 定义一个结构体类型的实例
	stu := Student{
		Name: "Alan",
		Age:  29,
	}
	testReflect(stu)
}
