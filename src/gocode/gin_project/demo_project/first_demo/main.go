// 声明当前包为main包
// 在Go语言中，main包是可执行程序的入口包，编译后会生成可执行文件
package main

// 导入所需依赖包
import (
	// 导入Gin框架核心包，Gin是一个高性能的HTTP web框架，用于快速构建RESTful API
	"github.com/gin-gonic/gin"
	// 导入标准库的net/http包，主要用于获取HTTP状态码常量（如200、404等）
	"net/http"
)

// main函数是程序的入口点，程序启动时会首先执行这里的代码
func main() {
	// 创建Gin路由器实例，使用gin.Default()方法
	// 该方法会自动注册两个默认中间件：
	// 1. gin.Logger()：用于记录请求日志（包含请求路径、状态码、处理时间等信息）
	// 2. gin.Recovery()：用于捕获程序中的panic异常，避免程序崩溃，并返回500 Internal Server Error响应
	router := gin.Default()

	// 注册一个GET请求的路由处理器
	// 第一个参数是路由路径（"/ping"），客户端访问该路径时会触发第二个参数的处理函数
	router.GET("/ping", func(context *gin.Context) {
		// context是gin.Context类型的实例，封装了HTTP请求和响应的所有信息，提供了便捷的处理方法

		// 使用context.JSON()方法返回JSON格式的响应
		// 第一个参数是HTTP状态码，这里使用http.StatusOK（即200，表示请求成功）
		// 第二个参数是响应数据，gin.H是Gin框架提供的便捷类型，本质是map[string]interface{}的别名
		// 这里返回一个包含"message"字段的JSON对象，值为"pong"
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// 启动HTTP服务器
	// router.Run()默认会在0.0.0.0:8080地址监听请求（兼容Windows系统的localhost:8080）
	// 也可以指定端口，例如router.Run(":8081")会监听8081端口
	// 该方法会阻塞当前进程，直到服务器被手动终止（如Ctrl+C）
	router.Run(":1111")
}

/*
代码功能说明：
这段代码基于 Gin 框架实现了一个最简单的 HTTP 服务，主要功能包括：
初始化带有日志和异常恢复功能的 Gin 路由器
注册/ping路径的 GET 请求处理器，返回{"message":"pong"}的 JSON 响应
在 8080 端口启动服务，接收并处理客户端请求
运行程序后，可通过 curl http://localhost:8080/ping 或 浏览器访问该地址，获取响应结果。
*/
