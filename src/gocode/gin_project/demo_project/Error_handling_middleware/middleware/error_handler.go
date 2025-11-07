// 包名：middleware（表示这是中间件模块）
package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ErrorHandler 捕获请求处理中的错误，返回统一格式的JSON响应
// 注意：函数名首字母大写（Exported），才能被其他包导入使用
func ErrorHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next() // 先执行后续的业务处理逻辑

		// 检查是否有错误被添加到上下文
		if len(context.Errors) > 0 {
			// 取最后一个错误（通常是最新产生的错误）
			err := context.Errors.Last().Err

			// 返回统一格式的错误响应
			context.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": err.Error(),
			})
		}
	}
}
