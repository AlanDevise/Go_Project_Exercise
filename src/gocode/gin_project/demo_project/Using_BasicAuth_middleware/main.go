package main

import (
	"github.com/gin-gonic/gin" // 导入Gin Web框架
	"net/http"                 // 导入HTTP标准库，用于状态码等
)

// 模拟一些私有数据，键为用户名，值为该用户的私密信息（邮箱、电话等）
var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func main() {
	// 初始化Gin默认路由器（包含日志和崩溃恢复中间件）
	router := gin.Default()

	// 创建一个需要基本认证的路由组
	// 使用gin.BasicAuth中间件，参数为账号密码映射（用户名:密码）
	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",    // 用户名foo，密码bar
		"austin": "1234",   // 用户名austin，密码1234
		"lena":   "hello2", // 用户名lena，密码hello2
		"manu":   "4321",   // 用户名manu，密码4321
	}))

	// 定义/admin/secrets接口，属于上面的认证路由组（需先认证才能访问）
	// 访问地址：localhost:8080/admin/secrets
	authorized.GET("/secrets", func(c *gin.Context) {
		// 从上下文获取当前认证通过的用户名（由BasicAuth中间件设置）
		user := c.MustGet(gin.AuthUserKey).(string)
		// 检查该用户是否有对应的私密信息
		if secret, ok := secrets[user]; ok {
			// 有则返回用户信息和私密数据
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			// 无则返回用户信息和无私密数据的提示
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})

	// 启动HTTP服务，监听在127.0.0.1:8080
	router.Run("127.0.0.1:8080")
}
