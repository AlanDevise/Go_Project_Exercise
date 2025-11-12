package main

// 导入必要的包：http包用于处理HTTP相关操作，gin框架用于快速构建Web服务
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化Gin路由器，使用Default()会自动添加日志(Logger)和崩溃恢复(Recovery)中间件
	router := gin.Default()

	// 静态目录映射：将本地"./assets"目录映射到URL路径"/assets"
	// 访问 http://localhost:8080/assets/xxx 会对应读取 ./assets/xxx 文件
	router.Static("/assets", "./assets")

	// 更灵活的静态文件系统映射：通过http.Dir指定本地目录"my_file_system"，映射到URL路径"/more_static"
	// 相比Static，StaticFS支持自定义文件系统实现（这里用默认的http.Dir）
	router.StaticFS("/more_static", http.Dir("my_file_system"))

	// 单个静态文件映射：将本地"./resources/favicon.ico"文件绑定到URL路径"/favicon.ico"
	// 浏览器访问网站时会自动请求/favicon.ico作为图标，这里直接指定对应文件
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")

	// 启动HTTP服务，监听0.0.0.0:8080（所有网络接口的8080端口）
	router.Run(":8080")
}
