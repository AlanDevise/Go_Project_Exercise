package main

import (
	"fmt"
)

// 定义动物的结构体
type Animal struct {
	Age    int
	Weight float32
}

// 给Animal绑定方法，喊叫
func (an *Animal) Shout() {
	fmt.Println("我可以大声喊叫")
}

// 给Animal绑定方法，自我展示
func (an *Animal) ShowInfo() {
	fmt.Printf("动物的年龄是：%v，动物的体重是：%v", an.Age, an.Weight)
}

type Cat struct {
	// 为了复用，体现继承，加入匿名结构体
	Animal
}

// 对Cat绑定特有方法
func (cat *Cat) scratch() {
	fmt.Println("我是小猫，我能挠人")
}

// 给Animal绑定方法，自我展示
func (cat *Cat) ShowInfo() {
	fmt.Printf("猫的年龄是：%v，猫的体重是：%v", cat.Age, cat.Weight)
}

func main() {
	// 创建Cat结构体对象
	cat := &Cat{}
	cat.Age = 3
	cat.Weight = 10.6
	cat.Shout()
	cat.scratch()
	cat.ShowInfo()
}
