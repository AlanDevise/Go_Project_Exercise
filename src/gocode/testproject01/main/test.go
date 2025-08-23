// 声明文件所在包，每个go文件均需要归属包
package main

// 导入依赖库，为了使用库中函数
import (
	"fmt"
	"strconv"
)

func main() { // 主函数

	name := "sadfasdf"
	age := 12
	ageStr := strconv.FormatInt(int64(age), 10)
	var age2, _ = strconv.ParseInt(ageStr, 10, 64)

	fmt.Println("Hello ",
		"world.")
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(ageStr)
	fmt.Println(age2)
	fmt.Printf("age2 的类型是%T, age2 = %v\n", age2, age2)

	// 使用取地址符
	fmt.Printf("age2的地址是：%p\n", &age2)
	var pointer *int64 = &age2
	fmt.Printf("pointer的值是：%p\n", pointer)
	fmt.Printf("pointer的地址是：%p\n", &pointer)
	fmt.Println("pointer指向的数据值是：", *pointer)

	var parameterOne int64 = +50
	parameterOne += 50
	fmt.Println(parameterOne)

	if 1 > 5 {
		fmt.Println("1小于5")
	}

	sum := 0
	alanString := "飘零"
	// go语言中只有for循环，没有其他的
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	fmt.Println(alanString)

	fmt.Println("><><><><><><><><")

	// 死循环实例
	// for {
	// 	fmt.Println("死掉了")
	// }

	for _, c := range "alandevise中文" {
		fmt.Printf("%c", c)
	}
	fmt.Println("\n><><><><><><><><")

	sum = 0
	// result := 0
	for i := 1; ; i++ {
		sum += i
		// if i == 100 {
		// 	result = sum
		// }
		if sum > 300 {
			break
		}
	}
	fmt.Println(sum)

	fmt.Println("\n><><><><><><><><")

}
