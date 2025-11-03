package main

import (
	"fmt"
	"net"
)

func main() {
	// 打印：
	fmt.Println("[Server] 服务端启动……")
	// 进行监听，需要指定服务器TCP协议，服务器端的IP+port
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("[ERROR] 服务端创建监听器失败:", err)
		return
	}
	for {
		// 等待客户端连接
		connection, errorMsg := listen.Accept()
		if errorMsg != nil {
			fmt.Println("[ERROR] 监听器监听失败：", errorMsg)
		} else {
			// 连接成功
			fmt.Printf("[INFO] 等待连接成功：conn:%v，接收到的客户端信息：%v\n", connection, connection.RemoteAddr().String())
		}

		// 准备一个协程，协程处理客户端服务的请求
		go process(connection)
	}
}

func process(connection net.Conn) {
	defer connection.Close()
	for {
		buffer := make([]byte, 1024)
		n, err := connection.Read(buffer)
		if err != nil {
			fmt.Printf("[INFO] 客户端断开连接: %v\n", err)
			return
		}

		msg := string(buffer[:n])
		fmt.Printf("[INFO] 收到消息: %s\n", msg)

		if msg == "quit" {
			fmt.Println("[INFO] 客户端请求退出")
			return // 使用 return 而不是 break
		}
	}
}
