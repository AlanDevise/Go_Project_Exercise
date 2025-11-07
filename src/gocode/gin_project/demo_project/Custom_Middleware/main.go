// 声明主包，表明这是一个可执行程序
package main

// 导入所需依赖包：gin框架用于构建Web服务、log用于日志输出、time用于时间处理
import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// Logger 自定义日志中间件，用于记录请求处理耗时和响应状态码
func Logger() gin.HandlerFunc {
	// 返回符合gin中间件类型的函数
	return func(context *gin.Context) {
		// 记录请求开始的时间
		nowTime := time.Now()

		// 在上下文对象中设置一个示例变量"example"，值为"12345"
		context.Set("example", "12345")

		// 执行后续的处理逻辑（如路由对应的处理函数）
		context.Next()

		// 计算请求处理的总耗时（从开始到当前的时间差）
		latency := time.Since(nowTime)
		// 打印请求耗时
		log.Print(latency)

		// 获取响应的状态码
		status := context.Writer.Status()
		// 打印响应状态码
		log.Println(status)
	}
}

func main() {
	// 创建一个新的Gin引擎实例（不含默认中间件）
	router := gin.New()
	// 注册自定义的Logger中间件，使所有请求都经过该中间件处理
	router.Use(Logger())

	// 定义一个GET请求的路由，路径为"/test"，对应的处理函数如下
	router.GET("/test", func(context *gin.Context) {
		// 从上下文获取"example"变量（必须存在，否则会触发panic），并转换为字符串类型
		example := context.MustGet("example").(string)
		// 打印获取到的变量值，预期输出"12345"
		log.Println(example)
	})

	// 启动Web服务，监听8080端口
	router.Run(":8080")
}
