package main

import "fmt"

func main() {

	var arr []int = []int{1, 2, 3, 4, 5}
	var b []int = make([]int, 10)
	// 拷贝
	copy(b, arr) // 将arr的元素拷贝到b中
	fmt.Println("Array:", arr)
	fmt.Println("Copied Slice:", b)

	// 修改切片的元素不会影响底层数组
	b[0] = 100
	fmt.Println("Modified Copied Slice:", b)
	fmt.Println("Array after modifying copied slice:", arr)
	fmt.Println("========================")

	// 所以copy是一个深拷贝过程吗
	fmt.Printf("Address of arr: %p\n", &arr)
	fmt.Printf("Address of b: %p\n", &b)
	fmt.Printf("Address of arr[0]: %p\n", &arr[0])
	fmt.Printf("Address of b[0]: %p\n", &b[0])

}
