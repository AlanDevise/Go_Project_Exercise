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
}

func main() {
	// 对基本数据类型进行反射
	// 定义一个基本数据类型
	var num int = 100
	testReflect(num)
}
