package main // 声明为可执行程序包

import (
	"log" // 引入日志包，用于打印日志信息

	"github.com/gin-gonic/gin" // 引入Gin框架，用于快速构建Web应用
)

// Person 定义结构体，用于接收HTTP请求中的查询参数
// form标签指定参数名，与URL中的查询参数对应
type Person struct {
	Name    string `form:"name"`    // 姓名，对应查询参数name
	Address string `form:"address"` // 地址，对应查询参数address
}

func main() {
	// 创建Gin默认路由引擎，默认包含日志和恢复中间件
	route := gin.Default()
	// 注册/testing路由，支持所有HTTP方法（GET/POST等），由startPage函数处理请求
	route.Any("/testing", startPage)
	// 启动HTTP服务，监听本地8080端口
	err := route.Run("127.0.0.1:8080")
	if err != nil {
		return
	}
}

// startPage 处理/testing路由的请求
func startPage(c *gin.Context) {
	var person Person // 声明Person实例，用于存储绑定的查询参数

	// 尝试将URL查询参数绑定到person结构体
	// ShouldBindQuery仅绑定查询参数（忽略请求体），绑定成功返回nil
	if c.ShouldBindQuery(&person) == nil {
		log.Println("====== 仅通过查询参数绑定数据 ======")
		log.Println("姓名:", person.Name)    // 打印绑定的姓名
		log.Println("地址:", person.Address) // 打印绑定的地址
	}

	// 向客户端返回200状态码和"Success"响应
	c.String(200, "Success")
}
