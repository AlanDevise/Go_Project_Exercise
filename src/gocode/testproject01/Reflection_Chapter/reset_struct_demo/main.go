package main // 声明主包，可执行程序必须位于main包

import (
	"fmt"     // 导入fmt包，用于格式化输入输出
	"reflect" // 导入reflect包，提供运行时反射能力，可动态操作变量的类型和值
)

// Student 定义学生结构体，包含姓名和年龄字段
type Student struct {
	Name string // 学生姓名
	Age  int    // 学生年龄
}

// CPrint 学生结构体的方法，用于打印学生信息
func (s Student) CPrint() {
	fmt.Println("调用到了CPrint()方法")
	fmt.Println("学生的名字是：", s.Name)
}

// AGetSum 学生结构体的方法，计算两个整数的和并返回
func (s Student) AGetSum(n1, n2 int) int {
	println("调用了AGetSum")
	return n1 + n2
}

// BSet 学生结构体的方法，设置学生的姓名和年龄（值接收者，修改的是副本）
func (s Student) BSet(name string, age int) {
	s.Name = name
	s.Age = age
}

func TestStudentReset(a any) {
	val := reflect.ValueOf(a)
	fmt.Println(val)
	n := val.Elem().NumField()
	fmt.Println(n)
	// 修改字段值
	val.Elem().Field(0).SetString("Corrine")
}

// TestStudentStruct 测试通过反射操作Student结构体的字段和方法
// 参数a为任意类型（此处实际传入Student实例）
func TestStudentStruct(a any) {
	// 将传入的变量a转换为reflect.Value类型，用于反射操作
	val := reflect.ValueOf(a)
	fmt.Println("反射获取的实例值：", val)

	// 获取结构体的字段数量
	n1 := val.NumField()
	fmt.Println("结构体字段数量：", n1)

	// 遍历结构体的所有字段并打印
	for i := range n1 {
		fmt.Printf("第%d个字段的值:%v\n", i, val.Field(i)) // Field(i)获取第i个字段的值
	}

	// 获取结构体的方法数量（按方法名ASCII排序）
	n2 := val.NumMethod()
	fmt.Println("结构体方法数量：", n2)

	// 调用索引为2的方法（CPrint，无参数）：方法按名排序为AGetSum(0)、BSet(1)、CPrint(2)
	val.Method(2).Call(nil) // Call(nil)表示无参数调用

	// 准备AGetSum方法的参数（10和20），需包装为reflect.Value切片
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10)) // 将int转换为reflect.Value
	params = append(params, reflect.ValueOf(20))

	// 调用索引为0的方法（AGetSum），传入参数并获取返回值
	result := val.Method(0).Call(params)
	fmt.Println("AGetSum方法的返回值为：", result[0].Int()) // 取出第一个返回值并转为int
}

func main() {
	// 创建Student结构体实例
	s := Student{
		Name: "Alan",
		Age:  29,
	}
	// 调用测试函数，通过反射操作s的字段和方法
	TestStudentStruct(s)
	TestStudentReset(&s)
	fmt.Println(s)
}
