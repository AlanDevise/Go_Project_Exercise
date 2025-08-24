package main

import "fmt"

// 主函数入口
func main() {
	var stuScore string = "80"
	switch stuScore {
	case "100":
		fmt.Println("A+")
	case "90":
		fmt.Println("A")
	case "80":
		fmt.Println("B")
		fallthrough
	case "70":
		fmt.Println("C")
	default:
		fmt.Println("D")
	}
}
