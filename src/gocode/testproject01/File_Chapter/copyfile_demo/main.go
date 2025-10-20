package main

import (
	"fmt"
	"os"
)

func main() {
	// 定义源文件
	file1Path := "/Users/alanzhang/Go_Project_Exercise/src/gocode/testproject01/File_Chapter/openfile_demo/Test.txt"
	// 定义目标文件
	file2Path := "/Users/alanzhang/Go_Project_Exercise/src/gocode/testproject01/File_Chapter/copyfile_demo/destination.txt"

	file1, error := os.ReadFile(file1Path)
	if error != nil {
		fmt.Printf("读取源文件异常：%v", error)
		return
	}
	// 写出文件
	writeError := os.WriteFile(file2Path, file1, 0666)
	if writeError != nil {
		fmt.Println("写出失败")
	}
}
