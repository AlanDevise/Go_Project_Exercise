package main

import (
	"context"
	"github.com/gin-gonic/gin" // Gin Web框架
	"log"                      // 日志包
	"net/http"                 // HTTP服务器包
	"os"                       // 操作系统功能包
	"os/signal"                // 信号处理包
	"syscall"                  // 系统调用包
	"time"                     // 时间处理包
)

func main() {
	// 初始化Gin默认路由器（包含日志和崩溃恢复中间件）
	router := gin.Default()

	// 定义根路径GET请求的处理函数
	// 模拟耗时操作（睡眠5秒），然后返回欢迎信息
	router.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	// 创建HTTP服务器实例，指定监听地址和处理器（使用Gin路由）
	srv := &http.Server{
		Addr:    "127.0.0.1:8080", // 监听8080端口
		Handler: router.Handler(), // 关联Gin路由处理器
	}

	// 启动goroutine运行HTTP服务器（非阻塞，避免主线程被阻塞）
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// 监听失败且非正常关闭时，记录致命错误
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 配置信号监听：等待中断信号（如Ctrl+C）以优雅关闭服务器
	quit := make(chan os.Signal, 1) // 缓冲通道，用于接收信号
	// 监听SIGINT（中断信号）和SIGTERM（终止信号）
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit // 阻塞等待信号到来
	log.Println("Shutdown Server ...")

	// 创建带5秒超时的上下文，用于控制关闭流程
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // 确保上下文最终被取消，释放资源

	// 优雅关闭服务器：等待现有请求处理完成（最长5秒）
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Server Shutdown:", err) // 关闭过程出错时记录
	}
	log.Println("Server exiting") // 服务器成功退出
}
