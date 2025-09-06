package main

import "fmt"

func main() {
	// Declare and initialize a one-dimensional array
	var intarr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("One-dimensional array:", intarr)

	// Declare and initialize a slice
	// 切出一个片段，索引：从1开始到3结束，但不包含3
	// 切片的长度是2，容量是4
	// 切片的底层数组是原始数组的一部分
	// 切片的元素类型与底层数组的元素类型相同
	// 切片的长度和容量可以通过len()和cap()函数获取
	// 切片可以通过append()函数动态增加元素
	// 切片可以通过copy()函数复制元素
	var slice []int = intarr[1:3]
	fmt.Println("Slice from array:", slice)

	// 输出切片元素个数
	fmt.Println("Length of slice:", len(slice))
	// 输出切片容量
	fmt.Println("Capacity of slice:", cap(slice))

	// 切片的容量大小是如何确定的？
	// 切片的容量是从切片的起始位置到底层数组的末尾位置的元素个数
	// 例如：intarr[1:3]，起始位置是索引1，底层数组的末尾位置是索引4
	// 所以容量是4 - 1 = 3
	// 如果切片的起始位置是索引0，底层数组的末尾位置是索引4
	// 那么容量就是4 - 0 = 4

	// 切片的地址
	fmt.Printf("Address of array: %p\n", &intarr)
	fmt.Printf("Address of slice: %p\n", &slice)

	// 切片由三个部分组成的结构体
	// type slice struct {
	// 	ptr *int    // 指向底层数组的指针
	// 	len int     // 切片的长度
	// 	cap int     // 切片的容量
	// }

	// 修改切片的元素会影响底层数组
	slice[0] = 100
	fmt.Println("Modified slice:", slice)
	fmt.Println("Array after modifying slice:", intarr)
}
