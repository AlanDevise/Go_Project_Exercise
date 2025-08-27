package main

import (
	"errors"
	"fmt"
)

func main() {
	error_msg := test()
	if error_msg != nil {
		fmt.Println("捕获到异常:", error_msg)
	}
	fmt.Println("上面的触发执行成功")
	fmt.Println("正常执行下面的逻辑")
}

func test() (error_msg error) {
	num1 := 10
	num2 := 0
	if num2 == 0 {
		// 抛出一个自定义的异常
		error_msg := errors.New("除数不能为0")
		return error_msg
	} else {
		result := num1 / num2
		fmt.Println("结果是:", result)
		// 正常返回nil
		return nil
	}
}
