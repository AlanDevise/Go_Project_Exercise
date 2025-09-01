package main

import "fmt"

// 主函数入口
func main() {
	fmt.Println("Hello, World!")
	// 给出五个学生的成绩，求出总和成绩，及成绩的平均数（初始化赋值）
	scores := [5]float64{90.5, 78.0, 88.5, 67.0}
	// 利用下标赋值
	scores[4] = 100.0 // 给第五个学生赋值
	// 除了使用这种方式赋值外，还可以使用下面这种方式
	// scores := [...]float64{90.5, 78.0, 88.5, 67.0, 100.0}

	var sum float64 = 0.0
	for _, score := range scores {
		sum += score
	}
	avg := sum / float64(len(scores))
	fmt.Printf("总成绩: %.2f, 平均成绩: %.2f\n", sum, avg)

	fmt.Println()

	// 第二个数组
	var arr2 [5]int64

	fmt.Printf("arr2 = %p\n", &arr2)
	fmt.Printf("arr2的首元素地址 = %p\n", &arr2[0])
	fmt.Printf("arr2的次元素地址 = %p\n", &arr2[1])
	fmt.Println("arr2 len =", len(arr2))

	var array3 = [...]int{1, 3, 5, 7, 9}
	fmt.Println("array3 =", array3)

	var array4 = [...]int{1: 2, 3: 4, 5: 6} // 下标1是2，下标3是4，下标5是6
	fmt.Println("array4 =", array4)
}
