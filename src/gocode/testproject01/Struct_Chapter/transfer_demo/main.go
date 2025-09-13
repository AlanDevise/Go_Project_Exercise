package main

import "fmt"

type Person struct {
	Age  int
	Name string
}

type Student struct {
	Age  int
	Name string
}

func main() {
	alan := &Person{Age: 29}
	alanStu := &Student{}

	alanStu = (*Student)(alan)

	fmt.Printf("Alan's age from Person struct: %d\n", alan.Age)
	fmt.Printf("Alan's age from Student struct: %d\n", alanStu.Age)
}
