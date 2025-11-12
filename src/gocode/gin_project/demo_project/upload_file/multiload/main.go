package main

import (
	"fmt"      // 用于格式化输出
	"log"      // 用于日志记录
	"net/http" // 提供HTTP状态码等常量

	"github.com/gin-gonic/gin" // 引入Gin框架，用于快速构建HTTP服务
)

func main() {
	// 创建Gin路由器实例，默认包含日志和恢复中间件(处理panic)
	router := gin.Default()

	// 设置多部分表单的内存限制(上传文件时临时存储的内存上限)
	// 默认是32MiB，这里设置为8MiB(1 << 20 是1MiB，8<<20是8MiB)
	// 它的作用是限制解析多部分表单（包含上传文件）时使用的内存总量。
	// 当处理文件上传时，Gin 会先将上传的数据临时存储在内存中，若数据总量（所有文件 + 表单字段的总数据）超过 MaxMultipartMemory 设定的值，
	// 超出的部分会被写入磁盘临时文件（而不是继续占用内存）
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	// 注册POST请求处理路由，路径为"/upload"
	router.POST("/upload", func(c *gin.Context) {
		// 获取多部分表单数据(包含上传的文件)
		form, _ := c.MultipartForm()
		// 从表单中提取name为"files"的文件列表
		files := form.File["files"]

		// 遍历所有上传的文件
		for _, file := range files {
			// 打印当前文件名到日志
			log.Println(file.Filename)

			// 将文件保存到本地"./files/"目录下，文件名保持原文件名
			c.SaveUploadedFile(file, "./files/"+file.Filename)
		}

		// 向客户端返回成功响应，包含上传的文件数量
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})

	// 启动HTTP服务，监听本地8080端口
	router.Run("127.0.0.1:8080")
}
