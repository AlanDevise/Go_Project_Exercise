package main

import "github.com/gin-gonic/gin" // 导入Gin框架，用于快速构建HTTP服务器

// loginEndpoint 处理登录请求，返回包含"login endpoint"的JSON响应
func loginEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "login endpoint",
	})
}

// submitEndpoint 处理提交请求，返回包含"submit endpoint"的JSON响应
func submitEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "submit endpoint",
	})
}

// readEndpoint 处理读取请求，返回包含"read endpoint"的JSON响应
func readEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "read endpoint",
	})
}

func main() {
	// 初始化默认的Gin路由器，默认包含日志和崩溃恢复中间件
	router := gin.Default()

	// 定义v1版本的API路由组
	{
		v1 := router.Group("/v1")          // 路由组前缀为/v1
		v1.POST("/login", loginEndpoint)   // 注册POST /v1/login端点，绑定登录处理函数
		v1.POST("/submit", submitEndpoint) // 注册POST /v1/submit端点，绑定提交处理函数
		v1.POST("/read", readEndpoint)     // 注册POST /v1/read端点，绑定读取处理函数
	}

	// 定义v2版本的API路由组
	{
		v2 := router.Group("/v2")          // 路由组前缀为/v2
		v2.POST("/login", loginEndpoint)   // 注册POST /v2/login端点，绑定登录处理函数
		v2.POST("/submit", submitEndpoint) // 注册POST /v2/submit端点，绑定提交处理函数
		v2.POST("/read", readEndpoint)     // 注册POST /v2/read端点，绑定读取处理函数
	}

	// 启动HTTP服务器，监听127.0.0.1:8080地址
	router.Run("127.0.0.1:8080")
}
