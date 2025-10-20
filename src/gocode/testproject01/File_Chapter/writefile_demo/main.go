package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 写入文件操作：
	// 打开文件
	file, err := os.OpenFile("/Users/alanzhang/Go_Project_Exercise/src/gocode/testproject01/File_Chapter/writefile_demo/Test.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)

	if err != nil {
		fmt.Printf("打开文件失败：%v", err)
		return
	}

	// 将文件关闭
	defer file.Close()

	// 写入
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("你好我是Alan\n")
	}

	// 流带缓冲区，刷新数据-》真正写入文件中
	writer.Flush()
}
