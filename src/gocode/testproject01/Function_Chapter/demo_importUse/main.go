package main

import (
	"alan_utils"
	"fmt"
)

func main() {
	fmt.Println("This is a demo for import and use.")
	// 函数在调用的时候需要定位到所在的包
	alanutils_test.GetConn()
}
