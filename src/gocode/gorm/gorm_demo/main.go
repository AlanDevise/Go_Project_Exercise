package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin" // Gin Web框架，用于快速构建HTTP服务
	"gorm.io/driver/postgres"  // PostgreSQL数据库驱动
	"gorm.io/gorm"             // GORM ORM框架，用于数据库操作
)

// User 定义用户数据模型，映射数据库表结构
type User struct {
	gorm.Model        // 嵌入GORM默认模型（包含ID、CreatedAt、UpdatedAt、DeletedAt字段）
	Name       string `gorm:"type:varchar(50);not null" json:"name"`      // 用户名，非空字符串
	Age        int    `gorm:"type:int;default:0" json:"age"`              // 年龄，默认0
	Email      string `gorm:"type:varchar(100);uniqueIndex" json:"email"` // 邮箱，唯一索引（不可重复）
}

var db *gorm.DB // 全局数据库连接实例

// initDB 初始化数据库连接，若User表不存在则创建
func initDB() error {
	// 数据库连接参数（主机、端口、用户、密码、数据库名等）
	dsn := "host=127.0.0.1 port=15432 user=postgres password=AlanDevise2025 dbname=AlanDevise sslmode=disable search_path=AlanDevise,public"

	// 建立PostgreSQL连接
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err // 连接失败返回错误
	}

	// 检查User表是否已存在
	migrator := db.Migrator()
	hasTable := migrator.HasTable(&User{})

	// 表不存在时，自动迁移创建表（根据User结构体定义）
	if !hasTable {
		if err := db.AutoMigrate(&User{}); err != nil {
			return err // 创建表失败返回错误
		}
	}

	return nil
}

// createUser 处理创建用户请求（POST /users）
func createUser(c *gin.Context) {
	var user User
	// 绑定请求体数据到User结构体（自动解析JSON/Form等格式）
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误：" + err.Error()})
		return
	}
	// 保存用户数据到数据库
	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建用户失败：" + err.Error()})
		return
	}
	// 返回创建成功响应
	c.JSON(http.StatusOK, gin.H{"message": "用户创建成功", "data": user})
}

// getUser 处理获取单个用户请求（GET /users/:id）
func getUser(c *gin.Context) {
	idStr := c.Param("id") // 获取URL路径中的id参数
	// 转换id为整数（GORM默认使用int类型主键）
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID格式错误"})
		return
	}
	var user User
	// 查询未被软删除的用户（deleted_at IS NULL）
	if err := db.Where("deleted_at IS NULL").First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败：" + err.Error()})
		}
		return
	}
	// 返回用户信息
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// listUsers 处理获取用户列表请求（GET /users）
func listUsers(c *gin.Context) {
	var users []User
	// 查询所有未被软删除的用户
	if err := db.Where("deleted_at IS NULL").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败：" + err.Error()})
		return
	}
	// 返回用户列表及总数
	c.JSON(http.StatusOK, gin.H{"total": len(users), "data": users})
}

// updateUser 处理更新用户请求（PUT /users/:id）
func updateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID格式错误"})
		return
	}
	// 先查询用户是否存在（未被软删除）
	var user User
	if err := db.Where("deleted_at IS NULL").First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}
	// 绑定更新数据
	var updateData User
	if err := c.ShouldBind(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误：" + err.Error()})
		return
	}
	// 执行更新操作（仅更新非零值字段）
	if err := db.Model(&user).Updates(updateData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新失败：" + err.Error()})
		return
	}
	// 返回更新成功响应
	c.JSON(http.StatusOK, gin.H{"message": "用户更新成功", "data": user})
}

// deleteUser 处理删除用户请求（DELETE /users/:id）
func deleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID格式错误"})
		return
	}
	// 执行软删除（仅更新deleted_at字段，不物理删除数据）
	if err := db.Delete(&User{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除失败：" + err.Error()})
		return
	}
	// 返回删除成功响应
	c.JSON(http.StatusOK, gin.H{"message": "用户删除成功"})
}

func main() {
	// 初始化数据库，失败则退出程序
	if err := initDB(); err != nil {
		panic("数据库初始化失败：" + err.Error())
	}

	// 创建Gin路由实例（默认包含日志和恢复中间件）
	router := gin.Default()
	// 定义用户相关API的路由组（路径前缀：/users）
	userGroup := router.Group("/users")
	{
		userGroup.POST("", createUser)       // 创建用户
		userGroup.GET("/:id", getUser)       // 获取单个用户
		userGroup.GET("", listUsers)         // 获取用户列表
		userGroup.PUT("/:id", updateUser)    // 更新用户
		userGroup.DELETE("/:id", deleteUser) // 删除用户
	}

	// 启动HTTP服务，监听本地8080端口
	router.Run("127.0.0.1:8080")
}
