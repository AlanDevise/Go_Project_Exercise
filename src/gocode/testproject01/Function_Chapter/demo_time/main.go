package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, World!")
	var a int = 10
	var b string = "Go"
	fmt.Println("a:", a, "b:", b)

	// 利用time包中的Now函数获取当前时间
	now := time.Now()
	fmt.Println("Current time:", now)
	fmt.Printf("Formatted time: %s\n", now.Format("2006-01-02 15:04:05"))
	fmt.Printf("Type: %T\n", now)
	fmt.Printf("Unix Timestamp: %d\n", now.Unix())
	fmt.Printf("Unix Nano Timestamp: %d\n", now.UnixNano())
	fmt.Printf("Year: %d, Month: %d, Day: %d\n", now.Year(), now.Month(), now.Day())
	fmt.Printf("Hour: %d, Minute: %d, Second: %d\n", now.Hour(), now.Minute(), now.Second())
	fmt.Printf("Weekday: %s\n", now.Weekday())
	fmt.Printf("Location: %s\n", now.Location())
	fmt.Printf("Is DST: %t\n", now.IsDST()) // 判断是否为夏令时
}
