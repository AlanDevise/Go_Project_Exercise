package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	// // 直接打开文件一次性读入到内存
	// fileContent, error := os.ReadFile("/Users/alanzhang/Go_Project_Exercise/src/gocode/testproject01/File_Chapter/openfile_demo/Test.txt")
	// if error != nil {
	// 	fmt.Println("文件打开异常，对应的错误为：", error)
	// } else {
	// 	fmt.Printf("文件=%v\n", string(fileContent))
	// }

	// 打开文件
	file, err := os.Open("/Users/alanzhang/Go_Project_Exercise/src/gocode/testproject01/File_Chapter/openfile_demo/Test.txt")

	if err != nil {
		fmt.Println("文件打开异常，对应的错误为：", err)
	}

	// 当函数退出时，让file关闭，防止内存泄漏
	defer file.Close()

	// 创建一个流
	reader := bufio.NewReader(file)
	// 读取流操作
	for {
		str, err := reader.ReadString('\n') // 读取到一个换行符就停止
		if err == io.EOF {                  // io.EOF 表示已经读取到文件结尾
			break
		}
		// 没有到文件结尾的话，正常输出文件内容
		fmt.Println(str)
	}
	// 结束
	fmt.Println("文件读取成功")
}
