package main

import (
	"demo_project/Error_handling_middleware/middleware" // 自定义错误处理中间件包
	"errors"                                            // 用于创建错误实例的标准库
	"github.com/gin-gonic/gin"                          // Gin Web框架，用于快速构建HTTP服务
	"net/http"                                          // 提供HTTP状态码、方法等标准定义
)

// main 函数是程序的入口点
func main() {
	// 初始化Gin默认路由器，默认包含日志中间件(gin.Logger())和崩溃恢复中间件(gin.Recovery())
	router := gin.Default()

	// 注册自定义错误处理中间件，用于统一捕获和处理请求过程中产生的错误
	router.Use(middleware.ErrorHandler())

	// 定义GET请求的/ok路由，模拟无错误的正常业务场景
	router.GET("/ok", func(context *gin.Context) {
		// 标记是否有错误发生，此处模拟"无错误"场景
		somethingWentWrong := false

		// 若模拟有错误，则将错误添加到上下文并提前返回
		if somethingWentWrong {
			context.Error(errors.New("something went wrong")) // 往上下文写入错误信息
			return
		}

		// 无错误时，返回200状态码和成功响应
		context.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Everything is fine!",
		})
	})

	// 定义GET请求的/error路由，模拟有错误的业务场景
	router.GET("/error", func(context *gin.Context) {
		// 标记是否有错误发生，此处模拟"有错误"场景
		somethingWentWrong := true

		// 模拟错误发生，将错误添加到上下文并提前返回
		if somethingWentWrong {
			context.Error(errors.New("something went wrong1")) // 往上下文写入错误信息
			context.Error(errors.New("something went wrong2")) // 往上下文写入错误信息
			context.Error(errors.New("something went wrong3")) // 往上下文写入错误信息
			return
		}

		// 以下代码在本场景中不会执行，仅作为无错误时的响应示例
		context.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Everything is fine!",
		})
	})

	// 启动HTTP服务，默认监听本地8080端口
	router.Run("127.0.0.1:8080")
}
