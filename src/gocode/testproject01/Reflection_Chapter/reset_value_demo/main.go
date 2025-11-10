package main // 声明主包，可执行程序必须位于main包

import (
	"fmt"     // 导入fmt包，用于格式化输入输出
	"reflect" // 导入reflect包，提供运行时反射能力，可操作变量的类型和值
)

// testReflect 利用反射机制修改传入变量的值
// 参数i为interface{}类型，可接收任意类型的参数
func testReflect(i any) {
	// 通过reflect.ValueOf获取传入参数的反射值对象
	reValue := reflect.ValueOf(i)
	// Elem()获取指针指向的底层元素（解引用），SetInt(40)将该元素的值修改为40
	// 注意：只有通过指针传递的变量，其反射值的底层元素才允许被修改（可设置）
	reValue.Elem().SetInt(40)
}

func main() {
	var num int = 100 // 定义int类型变量num，初始值为100
	// 传入num的指针&num，因为反射需要通过指针才能修改原始变量的值
	testReflect(&num)
	fmt.Println(num) // 打印修改后的num值，输出结果为40
}
