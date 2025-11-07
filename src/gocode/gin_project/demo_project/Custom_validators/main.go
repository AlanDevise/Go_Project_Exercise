package main

import (
	"fmt"
	"net/http" // 提供HTTP状态码等常量
	"time"     // 时间处理库

	"github.com/gin-gonic/gin"               // Gin框架，用于快速构建HTTP服务
	"github.com/gin-gonic/gin/binding"       // Gin的参数绑定工具
	"github.com/go-playground/validator/v10" // 数据验证库，Gin默认使用
)

// Booking 用于绑定和验证预订日期的数据结构
type Booking struct {
	// CheckIn 入住时间：从查询参数绑定，必填，需通过bookabledate验证，日期格式为2006-01-02
	CheckIn time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	// CheckOut 退房时间：从查询参数绑定，必填，需晚于CheckIn，需通过bookabledate验证，日期格式为2006-01-02
	// gtfield 是 greater than field 的缩写，表示当前字段的值必须大于指定字段的值。这里指定的字段是 CheckIn，即 CheckOut（退房时间）必须晚于 CheckIn（入住时间）
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn,bookabledate" time_format:"2006-01-02"`
}

// bookableDate 自定义验证函数：检查日期是否为可预订日期（不早于今天）
var bookableDate validator.Func = func(fieldLevel validator.FieldLevel) bool {
	// 将验证字段转换为time.Time类型
	date, ok := fieldLevel.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		// 若日期早于今天，则不可预订
		if today.After(date) {
			return false
		}
	}
	return true
}

func main() {
	// 初始化Gin默认路由引擎（包含日志和恢复中间件）
	router := gin.Default()

	// 注册自定义验证器bookableDate到Gin的验证引擎
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookableDate)
	}

	// 注册GET请求路由，路径为/bookable，处理函数为getBookable
	router.GET("/bookable", getBookable)
	// 启动HTTP服务，监听8085端口
	router.Run(":8085")
}

// getBookable 处理预订日期验证请求
func getBookable(context *gin.Context) {
	var oneBooking Booking
	// 绑定查询参数到Booking结构体，并进行验证
	// if err := context.ShouldBindWith(&oneBooking, binding.Query); err == nil {
	if err := context.ShouldBind(&oneBooking); err == nil {
		// 验证成功，返回成功信息
		context.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
	} else {
		// // 验证失败，返回错误信息
		// context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 提取具体错误信息
		var errMsg string
		if errs, ok := err.(validator.ValidationErrors); ok {
			for _, error := range errs {
				switch error.Tag() {
				case "required":
					errMsg += fmt.Sprintf("字段 %s 为必填项；", error.Field())
				case "bookabledate":
					errMsg += fmt.Sprintf("字段 %s 必须是今天或未来的日期；", error.Field())
				case "gtfield":
					errMsg += fmt.Sprintf("字段 %s 必须晚于 %s；", error.Field(), error.Param())
				case "time_format":
					errMsg += fmt.Sprintf("字段 %s 格式必须为 2006-01-02；", error.Field())
				}
			}
		} else {
			errMsg = err.Error()
		}
		context.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}
}
