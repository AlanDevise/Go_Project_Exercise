package main

import "fmt"

type Teacher struct {
	// 变量名字大写表示对外部包可见
	Name   string
	Age    int
	School string
}

func main() {
	teacherTeam := make(map[string]Teacher)

	teacherTeam["t1"] = Teacher{"Alice", 30, "High School"}
	teacherTeam["t2"] = Teacher{"Bob", 35, "Middle School"}
	// 键值对创建结构体更好，因为不会因为结构体的字段顺序变化而出错
	teacherTeam["t1"] = Teacher{Name: "Alice", Age: 30, School: "High School"}

	fmt.Println("Teacher Team:", teacherTeam)
	println("============================")

	alan := Teacher{"Alan", 29, "University of Nottingham"}
	fmt.Println("Alan's age:", alan.Age)
	println("============================")

	fmt.Printf("Teacher t1: %p\n", &alan)
	println("============================")

	var Corrine *Teacher = new(Teacher)
	Corrine.Name = "Corrine"
	Corrine.Age = 28
	Corrine.School = "Oxford Brookes University"
	fmt.Println("Corrine:", *Corrine)
	println("============================")

	var tom *Teacher = &Teacher{}
	tom.Name = "Tom"
	tom.Age = 40
	tom.School = "Cambridge University"
	fmt.Println("Tom:", *tom)
	println("============================")
}
