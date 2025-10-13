package main

import (
	"fmt"
)

type A struct {
	a int
	b string
}

type B struct {
	c int
	d string
}

type C struct {
	A
	B
	int
}

func main() {
	// 构建C结构体实例，干了一个多重继承
	c := C{A{10, "aaa"}, B{20, "ccc"}, 1000}
	fmt.Println(c.int)
}
