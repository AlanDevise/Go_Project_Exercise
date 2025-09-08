package main

import "fmt"

func main() {
	// 定义map对象
	var userInfo map[string]string
	// map在使用前必须使用make函数进行初始化，否则会报错panic: assignment to entry in nil map
	// 使用make函数创建map对象
	userInfo = make(map[string]string, 10)
	// 向map中添加键值对
	userInfo["name"] = "张三"
	userInfo["age"] = "18"
	userInfo["address"] = "北京"
	fmt.Println("userInfo=", userInfo)
	fmt.Println("userinfo size =", len(userInfo))

	anOtherMap := map[int]string{
		1: "Go",
		2: "Java",
		3: "Python",
	}
	fmt.Println("anOtherMap=", anOtherMap)
	fmt.Println("anOtherMap size =", len(anOtherMap))
	anOtherMap[4] = "JavaScript"
	fmt.Println("anOtherMap=", anOtherMap)
	fmt.Println("anOtherMap size =", len(anOtherMap))

	// 删除键值对
	delete(userInfo, "age")
	fmt.Println("after delete userInfo=", userInfo)
	// 通过key获取value
	name, exist := userInfo["name"]
	if exist {
		fmt.Println("name=", name)
	} else {
		fmt.Println("not found key name")
	}
	fmt.Println("===================================")

	// 遍历map
	for k, v := range userInfo {
		fmt.Println("key=", k, " value=", v)
	}
}
