package main

import "fmt"

// main 是程序的入口函数
func main() {
	// 调用 sum 函数，计算 10 和 20 的和，并将结果赋值给 result
	var ten = 10
	var twenty = 20
	fmt.Printf("The address of ten is %p\n", &ten)
	fmt.Printf("The address of twenty is %p\n", &twenty)
	var result, extra = sum_variable(ten, twenty)
	// 输出结果
	fmt.Printf("The result is:%d, and extra is %d.\n", result, extra)
	fmt.Println()
	fmt.Println("--- demo of swap two number ---")
	swapTwoNum(&ten, &twenty)
	fmt.Println("After swap, ten =", ten)
	fmt.Println("After swap, twenty =", twenty)
	fmt.Println()
	// 例如，可以将 sum_two 函数作为参数传递给 compute 函数
	var result_2 = compute(11, 21, sum_two)
	fmt.Println("The result of compute is:", result_2)
}

func swapTwoNum(pointerOfNum1 *int, pointerOfNum2 *int) {
	fmt.Printf("The address of pointerOfNum1 is %p\n", &pointerOfNum1)
	fmt.Printf("The address of pointerOfNum2 is %p\n", &pointerOfNum2)
	var temp = *pointerOfNum1
	*pointerOfNum1 = *pointerOfNum2
	*pointerOfNum2 = temp
}

// sum_two 接收两个 int 类型参数，返回它们的和
func sum_two(a int, b int) int {
	return a + b
}

// sum_three 接收三个 int 类型参数，返回它们的和
func sum_three(a int, b int, c int) int {
	return a + b + c
}

// sum_variable 接收可变数量的 int 类型参数，返回它们的和
func sum_variable(numbers ...int) (int, int) {
	total := 0
	for _, number := range numbers {
		fmt.Printf("number = %d and address of it is %p\n", number, &number)
		total += number
	}
	return total, 0
}

// golang不支持重载（函数名相同，形参不同）

// golang支持将函数作为参数进行传递
func compute(a int, b int, operation func(int, int) int) int {
	return operation(a, b)
}
