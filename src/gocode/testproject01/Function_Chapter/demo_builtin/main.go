package main

import "fmt"

func main() {
	// new分配内存，new函数的实参是一个类型而不是具体数值，new函数返回值是对应类型的指针 num: *int
	// num实际上就是一个指向int类型的指针，仅此而已
	num := new(int)
	fmt.Printf("%T\n", num)  // *int
	fmt.Println(*num)        // 0
	fmt.Printf("%v\n", &num) // 0x1400009a038
	fmt.Printf("%p\n", &num) // 0x1400009a038
}
