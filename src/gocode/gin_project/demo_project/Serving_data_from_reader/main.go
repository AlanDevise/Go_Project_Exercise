package main

import (
	"net/http"

	"github.com/gin-gonic/gin" // 导入gin框架，用于快速构建HTTP服务器
)

func main() {
	// 创建默认的gin路由引擎（包含日志和恢复中间件）
	router := gin.Default()

	// 注册GET请求处理函数，路径为"/someDataFromReader"
	router.GET("/someDataFromReader", func(c *gin.Context) {
		// 发起HTTP GET请求，获取GitHub上的gin框架logo图片
		response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
		// 若请求出错或响应状态码不是200（成功），返回503（服务不可用）状态码
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}

		// 获取响应体（图片数据读取器）
		reader := response.Body
		// 获取响应内容长度
		contentLength := response.ContentLength
		// 获取响应内容类型（图片的MIME类型）
		contentType := response.Header.Get("Content-Type")

		// 定义额外响应头，指定文件下载时的文件名
		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="gopher.png"`, // 触发客户端下载，文件名为gopher.png
		}

		// 将获取到的图片数据通过响应返回给客户端
		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})

	// 启动HTTP服务器，监听127.0.0.1:8080
	router.Run("127.0.0.1:8080")
}
