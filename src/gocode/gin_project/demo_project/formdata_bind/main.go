package main

import (
	// 导入gin框架，用于快速构建web服务
	"github.com/gin-gonic/gin"
)

// StructA 基础结构体，包含一个表单字段映射
type StructA struct {
	FieldA string `form:"field_a"` // 绑定表单字段field_a
}

// StructB 嵌套结构体示例，包含StructA实例和一个表单字段
type StructB struct {
	NestedStruct StructA // 嵌套StructA的实例
	FieldB       string  `form:"field_b"` // 绑定表单字段field_b
}

// StructC 嵌套结构体指针示例，包含StructA指针和一个表单字段
type StructC struct {
	NestedStructPointer *StructA // 嵌套StructA的指针
	FieldC              string   `form:"field_c"` // 绑定表单字段field_c
}

// StructD 嵌套匿名结构体示例，包含匿名结构体和一个表单字段
type StructD struct {
	NestedAnonyStruct struct { // 匿名结构体，无单独类型定义
		FieldX string `form:"field_x"` // 绑定表单字段field_x
	}
	FieldD string `form:"field_d"` // 绑定表单字段field_d
}

// GetDataB 处理/getb路由的请求
func GetDataB(context *gin.Context) {
	var b StructB
	// 将请求参数绑定到StructB实例
	context.Bind(&b)
	// 返回JSON响应，包含嵌套的StructA数据和FieldB
	context.JSON(200, gin.H{
		"a": b.NestedStruct,
		"b": b.FieldB,
	})
}

// GetDataC 处理/getc路由的请求
func GetDataC(context *gin.Context) {
	var b StructC
	// 将请求参数绑定到StructC实例
	context.Bind(&b)
	// 返回JSON响应，包含嵌套的StructA指针数据和FieldC
	context.JSON(200, gin.H{
		"a": b.NestedStructPointer,
		"c": b.FieldC,
	})
}

// GetDataD 处理/getd路由的请求
func GetDataD(context *gin.Context) {
	var b StructD
	// 将请求参数绑定到StructD实例
	context.Bind(&b)
	// 返回JSON响应，包含嵌套的匿名结构体数据和FieldD
	context.JSON(200, gin.H{
		"x": b.NestedAnonyStruct,
		"d": b.FieldD,
	})
}

// Student
type Student struct {
	// 由于Go的结构体字段首字母大写表示可导出，因此这里使用大写字段名
	Name string `form:"name"` // 绑定表单字段name
	Age  int    `form:"age"`  // 绑定表单字段age
}

func getStudent(context *gin.Context) {
	var student Student
	// 将请求参数绑定到Student实例
	context.Bind(&student)
	// 返回JSON响应，包含Student数据
	context.JSON(200, gin.H{
		"name": student.Name,
		"age":  student.Age,
	})
}

func main() {
	// 初始化gin默认路由器（包含日志和恢复中间件）
	router := gin.Default()
	// 注册GET路由及其处理函数
	router.GET("/getb", GetDataB)
	router.GET("/getc", GetDataC)
	router.GET("/getd", GetDataD)
	router.GET("/getStudent", getStudent)

	// 启动web服务器（默认监听8080端口）
	router.Run()
}
