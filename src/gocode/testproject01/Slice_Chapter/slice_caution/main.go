package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	var intarr [5]int = [5]int{1, 2, 3, 4, 5}

	var slice []int = intarr[1:4]
	fmt.Println("Empty slice:", slice)
	fmt.Println("Length of empty slice:", len(slice))
	fmt.Println("Capacity of empty slice:", cap(slice))
	fmt.Println("------------------------------------")
	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println(slice[2])

	// 修改切片的元素会影响底层数组
	slice[0] = 100
	fmt.Println("Modified slice:", slice)
	fmt.Println("Array after modifying slice:", intarr)
	fmt.Println("------------------------------------")

	// 切片能够继续切片
	slice2 := slice[1:3]
	fmt.Println("Slice2 from slice:", slice2)
	fmt.Println("Length of slice2:", len(slice2))
	fmt.Println("Capacity of slice2:", cap(slice2))
	fmt.Printf("Address of slice2: %p\n", &slice2)
	// 通过反射获取切片的 SliceHeader
	header := (*reflect.SliceHeader)(unsafe.Pointer(&slice2))
	// Data 字段即为底层数组的指针（uintptr 类型）
	fmt.Printf("底层数组指针: %#x\n", header.Data)
	fmt.Println("------------------------------------")

	// 切片能够动态增长，但是append方法会返回一个新的slice，并不是在原有的slice上进行追加
	// 但是这样是否有浪费空间的问题呢？
	slice2 = append(slice2, 200, 500)
	fmt.Println("Slice3 after append:", slice2)
	fmt.Println("Length of slice3:", len(slice2))
	fmt.Println("Capacity of slice3:", cap(slice2))
	fmt.Println("Array after appending to slice2:", intarr)
	fmt.Printf("Address of slice2: %p\n", &slice2)
	fmt.Println("------------------------------------")
	// 通过反射获取切片的 SliceHeader
	header2 := (*reflect.SliceHeader)(unsafe.Pointer(&slice2))
	// Data 字段即为底层数组的指针（uintptr 类型）
	fmt.Printf("底层数组指针: %#x\n", header2.Data)
}
