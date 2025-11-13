package main

import (
	"github.com/gin-gonic/gin" // 导入Gin框架包
	"io"                       // 导入io包，用于I/O操作
	"os"                       // 导入os包，用于文件操作
)

func main() {
	// 禁用控制台日志的颜色输出（日志写入文件时通常不需要颜色）
	gin.DisableConsoleColor()

	// 打开日志文件：若不存在则创建，若存在则追加写入（而非覆盖）
	// os.O_CREATE：文件不存在时创建
	// os.O_WRONLY：只写模式
	// os.O_APPEND：追加模式
	// 0644：文件权限（所有者读写，其他用户只读）
	f, err := os.OpenFile("gin.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		// 若文件打开失败，打印错误并退出
		panic(err)
	}
	// 设置Gin的默认日志输出目标为日志文件
	gin.DefaultWriter = io.MultiWriter(f)

	// 以下代码用于同时将日志输出到文件和控制台（当前被注释）
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// 创建Gin默认路由实例（默认包含日志和崩溃恢复中间件）
	router := gin.Default()
	// 注册GET请求路由：路径为/ping，处理函数返回状态码200和"pong"
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	// 启动HTTP服务，监听本地8080端口
	router.Run("127.0.0.1:8080")
}
