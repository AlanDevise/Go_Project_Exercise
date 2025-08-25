package main

import "fmt"

func main() {
	fmt.Println(add(30, 60))
	// defer的应用场景：
	// 1. 资源释放
	// 2. 文件关闭
	// 3. 解锁
	// 4. 打印日志
	// 5. 性能统计
	// 6. 异常处理
	// 方法执行完毕之后，自动关闭这个资源
}

func add(num1 int, num2 int) int {
	// 在golang中，程序遇到defer关键字，不会立即执行defer后面的语句，
	// 而是将defer后面的语句压入到一个栈中，然后继续执行函数后面的语句
	defer fmt.Println("num1 = ", num1)
	defer fmt.Println("num2 = ", num2)
	var sum int = num1 + num2
	fmt.Println("sum = ", sum)
	return sum
}
