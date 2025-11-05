package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/someJSON", func(context *gin.Context) {
		// responseData := map[string]any{
		// 	"lang": "GO语言",
		// 	"tag":  "<br>",
		// }
		// 上面的写法等价于
		responseData := gin.H{
			"lang": "GO语言",
			"tag":  "<br>",
		}

		// 使用 AsciiJSON 方法返回 JSON 响应，非 ASCII 字符将被转义
		// 如果客户端环境支持 UTF-8 编码，context.JSON() 更简洁；如果需要兼容低版本或编码支持差的环境，context.AsciiJSON() 更合适
		context.AsciiJSON(http.StatusOK, responseData)
		// context.JSON(http.StatusOK, responseData)
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	router.Run(":8080")
}
