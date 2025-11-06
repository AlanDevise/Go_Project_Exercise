package main

import (
	"github.com/gin-gonic/gin" // Gin Web框架，用于快速构建HTTP服务
	"log"                      // 日志包，用于输出日志信息
	"net/http"                 // HTTP状态码及相关常量定义
	"strconv"                  // 用于数值类型转换
	"time"                     // 时间处理包，用于时间类型的解析和处理
)

// Person 定义数据绑定结构体，用于接收和解析请求参数
type Person struct {
	Name     string    `form:"name"`                                       // 姓名，绑定请求中的"name"参数（表单或查询参数）
	Address  string    `form:"address"`                                    // 地址，绑定请求中的"address"参数
	Birthday time.Time `form:"birthday" time_format:"2006-01-02 15:04:05"` // 生日，绑定"birthday"参数
	// time_format指定解析格式（2006-01-02是Go的参考时间格式），time_utc:"1"表示解析为UTC时间
}

func main() {
	router := gin.Default()            // 初始化Gin默认路由器（包含日志和恢复中间件）
	router.GET("/testing", startPage)  // 注册GET请求路由：当访问/testing路径时，由startPage函数处理
	router.POST("/testing", startPage) // 注册GET请求路由：当访问/testing路径时，由startPage函数处理
	router.Run(":8085")                // 启动HTTP服务，监听8085端口
}

// startPage 处理/testing路径的GET请求，负责参数绑定、验证和响应
func startPage(context *gin.Context) {
	var person Person // 声明Person类型变量，用于接收绑定的请求参数

	// 数据绑定：根据请求方法自动选择绑定方式（GET用查询参数，POST根据Content-Type判断）
	if err := context.ShouldBind(&person); err != nil {
		// 绑定失败时，返回400状态码和错误信息
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 打印绑定成功的参数（日志输出）
	log.Println(person.Name)
	log.Println(person.Address)
	log.Println(person.Birthday)
	// 2. 手动转换为毫秒时间戳
	// milliTimestamp := person.Birthday.UnixNano() / 1e6
	milliTimestamp := person.Birthday.UnixMilli()
	log.Println("毫秒时间戳: " + strconv.FormatInt(milliTimestamp, 10))

	// 绑定成功，返回200状态和"Success"字符串
	context.String(200, "Success")
}
