package main

import (
	"fmt"
	// "testproject01/Function_Chapter/demo_importUse/alanutils_test"
)

func init() {
	fmt.Println("This is an init function in demo_importUse package.")
}

func main() {
	fmt.Println("This is a demo for import and use.")
	// 函数在调用的时候需要定位到所在的包，同一个包下不可定义重名的函数
	// alanutils_test.GetConn()
}
