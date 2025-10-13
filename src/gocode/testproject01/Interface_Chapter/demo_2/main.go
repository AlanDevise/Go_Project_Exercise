package main

import (
	"fmt"
)

type AInterface interface {
	a()
}

type BInterface interface {
	b()
}

type Stu struct {
	name string
	age  int
}

func (s Stu) a() {
	fmt.Println("aaaaaaa")
}

func (s Stu) b() {
	fmt.Println("bbbbbbb")
}

func main() {
	var s Stu
	var a AInterface = s
	var b BInterface = s
	a.a()
	b.b()
}
