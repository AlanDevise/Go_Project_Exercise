package main

import (
	// 导入Gin框架，用于快速构建Web应用
	"github.com/gin-gonic/gin"
)

// LoginForm 定义登录表单数据结构，用于接收和验证客户端提交的登录信息
type LoginForm struct {
	User     string `form:"user" binding:"required"`     // 用户名，form标签指定表单字段名，binding:"required"表示该字段为必填
	Password string `form:"password" binding:"required"` // 密码，同上，必填字段
}

func main() {
	// 初始化Gin路由器，使用默认配置（包含日志和恢复中间件）
	router := gin.Default()

	// 注册POST请求处理函数，路径为"/login"
	router.POST("/login", func(c *gin.Context) {
		// 声明LoginForm实例，用于存储绑定后的表单数据
		var form LoginForm

		// 自动绑定请求数据到form结构体（根据请求类型自动选择绑定方式，如表单、JSON等）
		// 同时会触发结构体中的binding验证规则（如必填校验）
		if c.ShouldBind(&form) == nil {
			// 验证用户名和密码是否正确（这里使用硬编码的"user"和"password"作为示例）
			if form.User == "user" && form.Password == "password" {
				// 验证通过，返回200状态码和登录成功信息
				c.JSON(200, gin.H{"status": "you are logged in"})
			} else {
				// 用户名或密码错误，返回401未授权状态码
				c.JSON(401, gin.H{"status": "unauthorized"})
			}
		}
	})

	// 启动HTTP服务器，监听本地8080端口
	router.Run("127.0.0.1:8080")
}
