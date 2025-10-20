package main

import (
	"io"
	"os"
)

func main() {
	// 定义源文件路径 - 要复制的原始文件
	srcPath := "/Users/alanzhang/Go_Project_Exercise/src/gocode/testproject01/File_Chapter/openfile_demo/Test.txt"
	// 定义目标文件路径 - 复制后生成的新文件
	dstPath := "/Users/alanzhang/Go_Project_Exercise/src/gocode/testproject01/File_Chapter/copyfile_demo2/destination.txt"

	// 打开源文件用于读取
	// os.Open() 以只读模式打开文件
	src, err := os.Open(srcPath)
	if err != nil {
		// 如果打开文件失败，立即终止程序并打印错误
		panic(err)
	}
	// 使用 defer 确保在函数退出前关闭源文件
	// 避免资源泄漏，无论函数如何返回都会执行
	defer src.Close()

	// 创建目标文件用于写入
	// os.Create() 会创建新文件，如果文件已存在则清空内容
	dst, err := os.Create(dstPath)
	if err != nil {
		// 如果创建文件失败，立即终止程序并打印错误
		panic(err)
	}
	// 使用 defer 确保在函数退出前关闭目标文件
	defer dst.Close()

	// 执行文件复制操作
	// io.Copy() 将源文件内容复制到目标文件
	// 返回值是复制的字节数和可能的错误
	_, err = io.Copy(dst, src)
	if err != nil {
		// 如果复制过程中出现错误，立即终止程序并打印错误
		panic(err)
	}

	// 程序执行到这里说明文件复制成功
	// 所有打开的文件会在函数返回时通过 defer 自动关闭
}
