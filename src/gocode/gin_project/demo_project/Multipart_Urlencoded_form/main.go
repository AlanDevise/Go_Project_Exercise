package main // 声明为可执行程序包

import (
	"github.com/gin-gonic/gin" // 导入Gin框架，用于快速构建HTTP服务
	"net/http"                 // 导入HTTP标准库，用于定义HTTP状态码等
)

func main() {
	// 创建一个默认的Gin路由引擎，默认包含日志和崩溃恢复中间件
	router := gin.Default()

	// 注册一个POST请求处理器，处理路径为"/form_post"的请求
	router.POST("/form_post", func(c *gin.Context) {
		// 从POST表单中获取"message"字段的值
		message := c.PostForm("message")
		// 从POST表单中获取"nick"字段的值，若不存在则使用默认值"anonymous"
		nick := c.DefaultPostForm("nick", "anonymous")

		// 以JSON格式返回响应，HTTP状态码为200(OK)
		// 响应内容包含状态、message字段值和nick字段值
		c.JSON(http.StatusOK, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	// 启动HTTP服务，监听本地8080端口
	router.Run("127.0.0.1:8080")
}
