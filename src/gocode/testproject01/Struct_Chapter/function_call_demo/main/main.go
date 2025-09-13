package main

import (
	"fmt"
	"testproject01/Struct_Chapter/function_call_demo/model"
)

func main() {
	student := model.NewStudent("Tom", 18)
	fmt.Printf("Student Name: %v, Age: %v\n", student.Name, student.Age)
	fmt.Println("==========================")
	fmt.Println("==========================")
}
