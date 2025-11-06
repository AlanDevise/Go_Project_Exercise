package main // 声明当前包为main，表明这是一个可执行程序

import (
	"github.com/gin-gonic/gin" // 导入Gin框架，用于快速构建HTTP服务器
	"net/http"                 // 导入HTTP标准库，提供HTTP服务器相关基础功能
	"time"                     // 导入时间库，用于设置超时时间
)

func main() { // 程序入口函数
	// 创建Gin默认路由引擎，默认包含日志(Logger)和恢复(Recovery)中间件
	router := gin.Default()

	// 通用 http.Server 配置
	httpServer := &http.Server{
		Addr:           ":8080",          // 监听地址（生产环境建议根据需求修改，如 ":80" 或 ":443"）
		Handler:        router,           // 处理器（Gin 路由引擎）
		ReadTimeout:    15 * time.Second, // 读取整个请求（包括 headers + body）的超时时间
		WriteTimeout:   20 * time.Second, // 写响应的超时时间（从接收到请求到发送完响应的总时间）
		IdleTimeout:    60 * time.Second, // 空闲连接超时（keep-alive 连接在无请求时的最大保持时间）
		MaxHeaderBytes: 4 << 20,          // 请求头最大字节数（4MB，比默认 1MB 更灵活，可根据业务调整）
		// TLSConfig: &tls.Config{                    // 若使用 HTTPS，需配置 TLS（可选）
		// 	MinVersion: tls.VersionTLS12,             // 禁用低版本 TLS（如 TLS 1.0/1.1）
		// 	CipherSuites: []uint16{                   // 只允许安全的加密套件
		// 		tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		// 		tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
		// 	},
		// 	PreferServerCipherSuites: true,
		// },
	}

	// 启动服务器，监听并处理请求
	httpServer.ListenAndServe()
}
