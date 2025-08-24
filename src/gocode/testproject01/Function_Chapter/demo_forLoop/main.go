package main

import "fmt"

func main() {
	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println("for loop i =", i)
	}

	// while loop
	j := 0
	for j < 5 {
		fmt.Println("while loop j =", j)
		j++
	}

	// infinite loop with break
	k := 0
	for {
		if k >= 5 {
			break
		}
		fmt.Println("infinite loop k =", k)
		k++
	}

	// for range loop over a slice
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("for range loop index = %d, value = %d\n", index, value)
	}

	// for range loop over a map
	person := map[string]int{"Alice": 30, "Bob": 25}
	for key, value := range person {
		fmt.Printf("for range loop key = %s, value = %d\n", key, value)
	}
}
