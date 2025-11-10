package main

import (
	"github.com/gin-gonic/gin" // 导入Gin框架，用于快速构建Web应用
	"net/http"                 // 导入HTTP包，提供HTTP状态码等基础功能
)

func main() {
	// 创建Gin默认路由引擎，默认包含日志(Logger)和恢复(Recovery)中间件
	router := gin.Default()

	// 定义GET请求路由：匹配路径为"/user/具体名称"的请求
	// 注意：仅匹配如"/user/john"这类路径，不匹配"/user/"或"/user"
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")                   // 通过Param方法获取路径参数"name"
		c.String(http.StatusOK, "Hello %s", name) // 返回字符串响应，状态码200
	})

	// 定义GET请求路由：匹配路径为"/user/具体名称/后续路径"的请求
	// 注意：可匹配如"/user/john/"或"/user/john/send"等路径，*action为通配符参数
	// 若没有其他路由匹配"/user/john"，会自动重定向到"/user/john/"
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")           // 获取路径参数"name"
		action := c.Param("action")       // 获取通配符参数"action"（包含后续所有路径）
		message := name + " is " + action // 组合消息内容
		c.String(http.StatusOK, message)  // 返回组合后的字符串响应，状态码200
	})

	// 启动HTTP服务器，监听本地8080端口
	router.Run("127.0.0.1:8080")
}
