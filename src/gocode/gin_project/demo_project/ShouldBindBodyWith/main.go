package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

// 定义两个测试用的结构体，用于绑定不同格式的请求体
type formA struct {
	// Gin 的默认绑定行为是 “宽松的”：只校验格式正确性，不强制要求字段存在（除非显式添加 binding:"required"）
	FieldA string `json:"field_a" xml:"field_a" binding:"required"`
}

type formB struct {
	FieldB string `json:"field_b" xml:"field_b" binding:"required"`
}

func main() {
	// 初始化gin引擎
	r := gin.Default()

	// 注册POST路由，使用我们定义的处理器
	r.POST("/test", SomeHandler)

	// 启动服务，监听8080端口
	r.Run(":8080")
}

// SomeHandler 演示ShouldBindBodyWith的复用绑定功能
func SomeHandler(c *gin.Context) {
	objA := formA{}
	objB := formB{}

	// 第一次绑定：尝试将请求体绑定到formA（JSON格式）
	if errA := c.ShouldBindBodyWith(&objA, binding.JSON); errA == nil {
		c.String(http.StatusOK, "匹配到formA格式：%+v", objA)
		return
	}

	// 第二次绑定：复用上下文的body，尝试绑定到formB（JSON格式）
	if errB := c.ShouldBindBodyWith(&objB, binding.JSON); errB == nil {
		c.String(http.StatusOK, "匹配到formB(JSON)格式：%+v", objB)
		return
	}

	// 第三次绑定：复用上下文的body，尝试绑定到formB（XML格式）
	if errB2 := c.ShouldBindBodyWith(&objB, binding.XML); errB2 == nil {
		c.String(http.StatusOK, "匹配到formB(XML)格式：%+v", objB)
		return
	}

	// 所有绑定都失败
	c.String(http.StatusNotFound, "未匹配到任何格式")
}
