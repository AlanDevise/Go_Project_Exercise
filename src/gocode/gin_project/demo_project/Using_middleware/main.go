package main

import (
	"github.com/gin-gonic/gin" // 导入Gin Web框架
)

func main() {
	// 创建一个没有默认中间件的路由器（默认中间件如日志、恢复等需手动添加）
	router := gin.New()

	// 全局中间件：日志中间件
	// 会将请求日志写入默认输出（如控制台），即使在release模式下也生效
	router.Use(gin.Logger())

	// 全局中间件：恢复中间件
	// 用于捕获程序中的panic，防止服务崩溃，并返回500状态码
	router.Use(gin.Recovery())

	// 路由级中间件：为特定路由添加自定义中间件
	// 访问"/benchmark"时，先执行MyBenchLogger()中间件，再执行benchEndpoint处理函数
	router.GET("/benchmark", MyBenchLogger(), benchEndpoint)

	// 创建授权路由组（基础路径为"/"）
	authorized := router.Group("/")
	// 为该路由组添加权限验证中间件（组内所有路由都会经过此验证）
	authorized.Use(AuthRequired())
	{
		// 授权组内的路由：需经过AuthRequired()验证才能访问
		authorized.POST("/login", loginEndpoint)   // 登录接口
		authorized.POST("/submit", submitEndpoint) // 提交接口
		authorized.POST("/read", readEndpoint)     // 读取接口

		// 嵌套路由组：在授权组内创建"testing"子组（路径为"/testing"）
		testing := authorized.Group("testing")
		// 嵌套组内的路由：需同时经过外层授权组和当前组的中间件（此处无额外中间件）
		testing.GET("/analytics", analyticsEndpoint) // 分析数据接口
	}

	// 启动服务，监听127.0.0.1:8080地址
	router.Run("127.0.0.1:8080")
}

// 以下为未实现的处理函数和中间件（仅作为示例声明）
func MyBenchLogger() gin.HandlerFunc   { return nil } // 自定义基准测试日志中间件
func benchEndpoint(c *gin.Context)     {}             // "/benchmark"路由处理函数
func AuthRequired() gin.HandlerFunc    { return nil } // 权限验证中间件
func loginEndpoint(c *gin.Context)     {}             // "/login"路由处理函数
func submitEndpoint(c *gin.Context)    {}             // "/submit"路由处理函数
func readEndpoint(c *gin.Context)      {}             // "/read"路由处理函数
func analyticsEndpoint(c *gin.Context) {}             // "/testing/analytics"路由处理函数
