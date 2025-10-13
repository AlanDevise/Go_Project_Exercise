package main

import (
	"fmt"
)

type SayHello interface {
	// 声明没有实现的方法
	sayHello()
}

// 接口的实现
// 中国人：
type Chinese struct {
	name string
	age  int
}

func (person Chinese) sayHello() {
	fmt.Println("你好")
}

func (person Chinese) niuYangge() {
	fmt.Println("在扭秧歌")
}

type American struct {
	name string
	age  int
}

func (person American) sayHello() {
	fmt.Println("Hi")
}

func (person American) doPopping() {
	fmt.Println("Do poping")
}

// 定义一个函数，专门用来各国人来打招呼，接收具备SayHello接口能力的变量
func greet(say SayHello) {
	say.sayHello()
	// 断言的写法
	// 看s是否能够专程Chinese并且赋值给ch变量// 类型断言：检查s是否为Chinese类型
	// 使用类型分支替代多个独立的类型断言
	switch person := say.(type) { // 这里的type是关键字，相当于say.(type)是一个固定写法
	case Chinese:
		person.niuYangge()
	case American:
		person.doPopping()
	// 可以继续添加其他实现了SayHello接口的类型
	// case Japanese:
	//     v.bow()
	default:
		// 可选：处理未匹配的类型
		fmt.Println("未知国籍的人打招呼")
	}
}

func main() {
	// 声明一个Chinese对象
	c := Chinese{}
	// 声明一个American对象
	a := American{}
	greet(c)
	greet(a)
}
