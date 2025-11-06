package main

import (
	"fmt"
	// 导入gin框架，用于快速构建Web应用
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建Gin默认路由器，默认包含日志和恢复中间件(处理panic)
	router := gin.Default()

	// 注册GET请求处理路由，路径为"/cookie"
	router.GET("/cookie", func(context *gin.Context) {
		// 尝试从请求中获取名为"gin_cookie"的Cookie
		cookie, err := context.Cookie("gin_cookie")

		// 如果获取Cookie失败(如首次访问无此Cookie)
		if err != nil {
			// 临时设置cookie值为"NotSet"
			cookie = "NotSet"
			// 向客户端设置Cookie：名称"gin_cookie"，值"test"，过期时间3600秒，
			// 路径"/"，域名"localhost"，非HTTPS模式，启用HTTPOnly(防止JS访问)
			context.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}

		// 在服务器端打印当前Cookie的值
		fmt.Printf("Cookie value: %s \n", cookie)
	})

	// 启动HTTP服务器，默认监听8080端口
	router.Run()
}
