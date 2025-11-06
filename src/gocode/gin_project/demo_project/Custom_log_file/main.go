package main // 声明为可执行程序包

import (
	"fmt"                      // 用于格式化字符串输出
	"github.com/gin-gonic/gin" // 导入Gin框架（轻量级HTTP web框架）
	"time"                     // 用于时间处理
)

func main() {
	// 创建一个新的Gin路由器（不含默认中间件，需手动添加）
	router := gin.New()

	// 自定义日志格式中间件：将请求日志输出到默认输出流（stdout）
	// 格式包含：客户端IP、时间戳、请求方法、路径、协议、状态码、处理耗时、用户代理、错误信息
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,                       // 客户端IP地址
			param.TimeStamp.Format(time.RFC1123), // 时间戳（RFC1123格式）
			param.Method,                         // HTTP请求方法（如GET、POST）
			param.Path,                           // 请求路径
			param.Request.Proto,                  // HTTP协议版本（如HTTP/1.1）
			param.StatusCode,                     // 响应状态码
			param.Latency,                        // 请求处理耗时
			param.Request.UserAgent(),            // 客户端用户代理（浏览器/工具标识）
			param.ErrorMessage,                   // 错误信息（无错误则为空）
		)
	}))

	// 添加错误恢复中间件：当处理请求时发生panic，自动恢复并返回500错误，避免程序崩溃
	router.Use(gin.Recovery())

	// 定义GET请求路由：路径为/ping，处理函数返回200状态码和"pong"字符串
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// 启动HTTP服务器，监听8080端口
	router.Run(":8080")
}
