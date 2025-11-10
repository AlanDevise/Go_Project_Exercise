package main

import (
	"github.com/gin-gonic/gin" // 导入Gin框架，用于快速构建Web应用
	"log"                      // 导入log包，用于日志输出
	"time"                     // 导入time包，用于时间相关操作（此处模拟耗时任务）
)

func main() {
	// 创建Gin默认路由引擎，默认包含日志(Logger)和恢复(Recovery)中间件
	router := gin.Default()

	// 注册GET请求处理接口：/long_async（异步处理长时间任务）
	router.GET("/long_async", func(context *gin.Context) {
		// 在 Go 的 Gin 框架中，在协程（goroutine）里使用 gin.Context 时需要用副本（context.Copy()），
		// 核心原因是 gin.Context 并非并发安全的，且其生命周期与请求绑定。
		cCp := context.Copy()
		// 启动goroutine异步处理长时间任务（不阻塞当前请求响应）
		go func() {
			// 模拟耗时5秒的任务
			time.Sleep(5 * time.Second)

			// 使用复制的上下文获取请求路径并打印日志（注意必须用副本）
			log.Println("Done! in path " + cCp.Request.URL.Path)
		}()
	})

	// 注册GET请求处理接口：/long_sync（同步处理长时间任务）
	router.GET("/long_sync", func(context *gin.Context) {
		// 模拟耗时5秒的任务（同步处理，会阻塞当前请求直到任务完成）
		time.Sleep(5 * time.Second)

		// 因未使用goroutine，可直接使用原上下文，打印请求路径日志
		log.Println("Done! in path " + context.Request.URL.Path)
	})

	// 启动服务器，监听127.0.0.1:8080地址
	router.Run("127.0.0.1:8080")
}
