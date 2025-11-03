package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// 打印：
	fmt.Println("[Client] 客户端启动……")
	// 调用dial函数，需要指定tcp协议，需要指定服务器的IP+port
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("[ERROR] 客户端连接失败:", err)
		return
	}
	fmt.Println("[INFO] 连接成功，conn:", conn)

	for {
		// 通过客户端发送单行数据，然后退出
		reader := bufio.NewReader(os.Stdin)
		// 从终端读取一行用户输入的信息，以\n作为结尾
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("[ERROR] 终端输入失败：err", err)
		} else {
			// 去除首尾空白字符
			str = strings.TrimSpace(str)

			n, err := conn.Write([]byte(str))
			if err != nil {
				fmt.Println("[ERROR] 发送消息失败：err", err)
			} else {
				fmt.Printf("[DEBUG] 发送成功，共%d个字节\n", n)
			}
			if str == "quit" {
				break
			}
		}
	}

}
