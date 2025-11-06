package main

import (
	"github.com/gin-gonic/gin" // 导入Gin框架，用于快速构建Web应用
	"net/http"                 // 导入HTTP包，用于HTTP状态码和相关功能
)

// Login 定义登录请求的数据结构，用于绑定不同格式的请求数据
// 标签分别指定了form、json、xml格式下的字段名，binding:"required"表示该字段为必填项
type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`            // 对应请求中的字段名是 "user"
	Password string `form:"password" json:"password" xml:"password" binding:"required"` // 对应请求中的字段名是 "password"
}

func main() {
	// 初始化Gin路由器，使用默认中间件（包含日志和恢复中间件）
	router := gin.Default()

	// 定义POST接口/loginJSON，用于处理JSON格式的登录请求
	// 示例请求数据：{"user": "manu", "password": "123"}
	router.POST("/loginJSON", func(c *gin.Context) {
		var json Login // 声明用于接收JSON数据的变量
		// 将请求体中的JSON数据绑定到json变量，失败则返回400错误
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 验证用户名和密码是否正确（硬编码为manu和123）
		if json.User != "manu" || json.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"}) // 验证失败返回401
			return
		}

		// 验证成功返回200和登录成功信息
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// 定义POST接口/loginXML，用于处理XML格式的登录请求
	// 示例请求数据：
	// <?xml version="1.0" encoding="UTF-8"?>
	// <root>
	//   <user>manu</user>
	//   <password>123</password>
	// </root>
	router.POST("/loginXML", func(c *gin.Context) {
		var xml Login // 声明用于接收XML数据的变量
		// 将请求体中的XML数据绑定到xml变量，失败则返回400错误
		if err := c.ShouldBindXML(&xml); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 验证用户名和密码是否正确
		if xml.User != "manu" || xml.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		// 验证成功返回登录成功信息
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// 定义POST接口/loginForm，用于处理HTML表单格式的登录请求
	// 示例请求数据：user=manu&password=123（form-data或x-www-form-urlencoded）
	router.POST("/loginForm", func(c *gin.Context) {
		var form Login // 声明用于接收表单数据的变量
		// 根据请求的Content-Type自动选择绑定器（表单格式），失败返回400错误
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 验证用户名和密码是否正确
		if form.User != "manu" || form.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		// 验证成功返回登录成功信息
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// 启动HTTP服务器，监听0.0.0.0:8080端口
	router.Run(":8080")
}
