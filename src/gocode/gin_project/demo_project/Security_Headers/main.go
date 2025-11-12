package main

import (
	"net/http"

	"github.com/gin-gonic/gin" // 导入gin框架，用于快速开发Web服务
)

func main() {
	// 初始化gin引擎，默认包含日志中间件和panic恢复中间件
	r := gin.Default()

	// 定义预期的主机头值，用于验证请求合法性
	expectedHost := "127.0.0.1:8080"

	// 注册全局中间件：验证Host头并设置安全响应头
	r.Use(func(c *gin.Context) {
		// 验证请求的Host头是否匹配预期值，防止Host头攻击
		if c.Request.Host != expectedHost {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid host header"})
			return
		}

		// 设置安全相关的响应头，增强Web服务安全性
		c.Header("X-Frame-Options", "DENY")                                                                                                                                    // 禁止页面被嵌入iframe，防点击劫持
		c.Header("Content-Security-Policy", "default-src 'self'; connect-src *; font-src *; script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';") // 限制资源加载来源，防XSS等
		c.Header("X-XSS-Protection", "1; mode=block")                                                                                                                          // 启用浏览器XSS过滤器，检测到攻击时阻止加载
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")                                                                                           // 强制使用HTTPS，防HTTP降级攻击
		c.Header("Referrer-Policy", "strict-origin")                                                                                                                           // 控制Referrer信息发送规则，保护隐私
		c.Header("X-Content-Type-Options", "nosniff")                                                                                                                          // 禁止浏览器猜测内容类型，防MIME混淆攻击
		c.Header("Permissions-Policy", "geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=(),magnetometer=(),gyroscope=(),fullscreen=(self),payment=()")                 // 限制敏感权限使用

		c.Next() // 继续处理后续请求
	})

	// 注册GET类型路由/ping，用于测试服务可用性
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong", // 响应"pong"表示服务正常
		})
	})

	// 启动Web服务，监听本地127.0.0.1:8080端口
	r.Run("127.0.0.1:8080")
}
