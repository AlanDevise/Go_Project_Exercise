package model

type student struct {
	Name string
	Age  int
}

// 工厂模式
func NewStudent(name string, age int) student {
	return student{
		Name: name,
		Age:  age,
	}
}
