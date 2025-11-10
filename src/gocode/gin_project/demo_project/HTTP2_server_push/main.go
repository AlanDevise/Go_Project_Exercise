package main

import (
	"github.com/gin-gonic/gin" // 导入Gin Web框架
	"html/template"            // 导入模板处理包
	"log"                      // 导入日志包
)

// 定义HTML模板，包含页面结构及引用的静态资源(app.js)
var html = template.Must(template.New("https").Parse(`
<html>
<head>
  <title>Https Test</title>
  <script src="/assets/app.js"></script>
</head>
<body>
  <h1 style="color:red;">Welcome, Ginner!</h1>
</body>
</html>
`))

func main() {
	// 初始化Gin默认路由器（包含日志和崩溃恢复中间件）
	router := gin.Default()

	// 注册静态文件服务：将URL路径"/assets"映射到本地"./assets"目录
	router.Static("/assets", "./assets")

	// 设置Gin使用的HTML模板引擎
	router.SetHTMLTemplate(html)

	// 注册根路径("/")的GET请求处理器
	router.GET("/", func(c *gin.Context) {
		// 检查当前连接是否支持HTTP/2服务器推送
		if pusher := c.Writer.Pusher(); pusher != nil {
			// 推送静态资源app.js到客户端（减少客户端二次请求）
			if err := pusher.Push("/assets/app.js", nil); err != nil {
				log.Printf("推送资源失败: %v", err) // 推送失败时记录日志
			}
		}
		// 渲染HTML模板，状态码200，使用名为"https"的模板，传递数据
		c.HTML(200, "https", gin.H{
			"status": "success",
		})
	})

	// 启动HTTP服务器，监听127.0.0.1:8080
	// 注：若需启动HTTPS，可使用下面注释的RunTLS方法（需提供证书和密钥）
	// router.RunTLS(":8080", "./testdata/server.pem", "./testdata/server.key")
	router.Run("127.0.0.1:8080")
}
