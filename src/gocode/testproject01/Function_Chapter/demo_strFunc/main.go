package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str1 := "golang你好"
	fmt.Println("str1的长度为 = ", len(str1)) // 这里的长度是按照字节计算的，一个汉字3个字节
	fmt.Println()

	// 对字符串进行遍历
	for i := 0; i < len(str1); i++ {
		fmt.Printf("str1[%d] = %c\n", i, str1[i]) // 会出现乱码，因为是按照字节进行遍历的
	}
	fmt.Println()

	for index, value := range str1 { // 按照rune(字符)遍历，不会出现乱码
		fmt.Printf("index = %d, value = %c\n", index, value)
	}
	fmt.Println()

	// 转成切片进行遍历
	str2 := []rune(str1)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("str2[%d] = %c\n", i, str2[i])
	}
	fmt.Println()

	// 字符串转正数
	// strconv.Atoi("123")    // 123, nil
	// strconv.Atoi("123abc") // 0, error
	num, err := strconv.Atoi("123abc")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转换成功 num = ", num)
	}
	fmt.Println()

	// 正数转字符串
	str3 := strconv.Itoa(12345)
	fmt.Printf("str3的类型是 %T, str3 = %q\n", str3, str3)
	fmt.Println()

	// 统计一个字符串中有多少个指定的子串
	str4 := "asdfasdfasfd"
	substr := "as"
	// count := 0
	// for i := 0; i < len(str4)-len(substr)+1; i++ {
	// 	if str4[i:i+len(substr)] == substr {
	// 		count++
	// 	}
	// }
	// fmt.Printf("在字符串 %q 中, 子串 %q 出现的次数 = %d\n", str4, substr, count)
	count := strings.Count(str4, substr)
	fmt.Printf("在字符串 %q 中, 子串 %q 出现的次数 = %d\n", str4, substr, count)
	fmt.Println()

	// 不区分大小写进行比较
	result := strings.EqualFold("abc", "AbC")
	fmt.Println("result = ", result)
	fmt.Println()

	// 区分大小写进行比较
	fmt.Println("abc" == "AbC")
	fmt.Println()

	// 返回子串在字符串中第一次出现的索引位置, 如果没有返回-1
	index := strings.Index("asdfsadf", "e") // -1
	fmt.Println("index = ", index)
	fmt.Println()

}
