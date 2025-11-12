package main

import (
	"github.com/gin-gonic/gin"   // Gin框架，用于快速构建HTTP服务器
	"golang.org/x/sync/errgroup" // 用于管理并发goroutine，支持错误收集
	"log"                        // 日志输出包
	"net/http"                   // 标准库HTTP服务器相关功能
	"time"                       // 时间处理包
)

// 声明errgroup实例，用于管理两个服务器的goroutine
var g errgroup.Group

// 定义第一个服务器的路由处理器
func router01() http.Handler {
	e := gin.New()        // 创建Gin引擎实例（无默认中间件）
	e.Use(gin.Recovery()) // 使用Recovery中间件，用于捕获panic并返回500响应
	// 注册GET "/"路由，处理函数返回JSON响应
	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{ // gin.H是map[string]interface{}的别名，用于构建JSON数据
				"code":    http.StatusOK,
				"message": "Welcome server 01",
			},
		)
	})
	return e // 返回构建好的路由处理器
}

// 定义第二个服务器的路由处理器（结构与router01类似，响应信息不同）
func router02() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":    http.StatusOK,
				"message": "Welcome server 02",
			},
		)
	})
	return e
}

func main() {
	// 配置第一个HTTP服务器：监听8080端口，使用router01作为处理器
	server01 := &http.Server{
		Addr:         ":8080",          // 监听地址
		Handler:      router01(),       // 路由处理器
		ReadTimeout:  5 * time.Second,  // 读取请求超时时间
		WriteTimeout: 10 * time.Second, // 写入响应超时时间
	}

	// 配置第二个HTTP服务器：监听8081端口，使用router02作为处理器
	server02 := &http.Server{
		Addr:         ":8081",
		Handler:      router02(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// 通过errgroup启动第一个服务器的goroutine
	g.Go(func() error {
		return server01.ListenAndServe() // 启动服务器，监听并处理请求
	})

	// 通过errgroup启动第二个服务器的goroutine
	g.Go(func() error {
		return server02.ListenAndServe()
	})

	// 等待所有goroutine完成，若有错误则日志输出并退出
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
