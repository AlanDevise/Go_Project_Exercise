package main

import "github.com/gin-gonic/gin" // 导入Gin框架，用于快速构建HTTP服务

// Person 定义用于绑定URI参数的结构体
type Person struct {
	ID   string `uri:"id" binding:"required,uuid"` // 绑定URI中的"id"参数，验证规则：必填且为UUID格式
	Name string `uri:"name" binding:"required"`    // 绑定URI中的"name"参数，验证规则：必填
}

func main() {
	route := gin.Default() // 初始化Gin默认路由引擎（包含日志和恢复中间件）

	// route.GET("/thinkerou/987fbc97-4bed-5078-9f07-9141ba07c9f3", func(c *gin.Context) {
	// 	c.String(200, "hello, thinkerou")
	// })

	// 优先级顺序为 静态路由 > 参数化路由 > 通配符路由，嵌套路由的优先级由其内部具体路由类型决定
	// 注册GET请求路由，路径为"/:name/:id"，其中:name和:id为URI参数
	route.GET("/:name/:id", func(c *gin.Context) {
		var person Person // 声明Person实例用于接收绑定的URI参数

		// 将URI参数绑定到person结构体，并进行验证
		if err := c.ShouldBindUri(&person); err != nil {
			// 绑定/验证失败时，返回400错误和错误信息
			c.JSON(400, gin.H{"msg": err.Error()})
			return
		}

		// 绑定成功，返回200状态和解析后的参数
		c.JSON(200, gin.H{"name": person.Name, "uuid": person.ID})
	})

	route.Run(":8088") // 启动HTTP服务，监听8088端口
}
