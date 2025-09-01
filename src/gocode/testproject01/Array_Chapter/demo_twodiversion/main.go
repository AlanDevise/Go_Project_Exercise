package main

import "fmt"

func main() {
	var arr [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("arr[%d][%d]=%d ", i, j, arr[i][j])
		}
		fmt.Println()
	}

	fmt.Println()

	// 使用range方式
	for i, row := range arr {
		for j, val := range row {
			fmt.Printf("arr[%d][%d]=%d ", i, j, val)
		}
		fmt.Println()
	}

	fmt.Println()
}
