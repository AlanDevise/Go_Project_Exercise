package main // 声明为可执行程序包

import (
	"fmt"                      // 导入格式化输入输出包
	"github.com/gin-gonic/gin" // 导入Gin Web框架
)

func main() {
	// 创建默认的Gin路由器，包含日志和崩溃恢复中间件
	router := gin.Default()

	// 注册一个POST请求处理路由，路径为"/post"
	router.POST("/post", func(c *gin.Context) {
		// 从URL查询参数中获取键为"ids"的多个值，解析为map[string][]string
		ids := c.QueryMap("ids")
		// 从POST表单数据中获取键为"names"的多个值，解析为map[string][]string
		names := c.PostFormMap("names")

		// 打印获取到的ids和names数据
		fmt.Printf("ids: %v; names: %v", ids, names)
	})

	// 启动HTTP服务器，监听本地8080端口
	router.Run("127.0.0.1:8080")
}
