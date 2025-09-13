package main

import "fmt"

type A struct {
	name string
	age  int
}

func (a *A) PrintInfo() {
	a.name = " (modified)"
	fmt.Printf("name=%v, age=%v\n", a.name, a.age)
}

// 这里类似一个toString方法
func (a *A) String() string {
	return fmt.Sprintf("A{name=%v, age=%v}", a.name, a.age)
}

func main() {
	var a A
	a.name = "Alice"
	// 这里有个语法糖，可以直接用a调用PrintInfo方法
	a.PrintInfo()
	fmt.Printf("After PrintInfo call: name=%v, age=%v\n", a.name, a.age)
	fmt.Println("=====================")

	fmt.Println(&a)
}
