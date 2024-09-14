package main

import "fmt"

func main() {
	students := make(map[string]int)
	students["Sara"] = 10
	students["Evan"] = 8
	students["Henna"] = 7

	val, isExist := students["Evan"]

	fmt.Println("Evan got :", val)
	fmt.Println("Evan is here :", isExist)

	// ------------------------

	numbers := map[string]int{
		"Sara":  10,
		"Evan":  9,
		"Henna": 8,
	}

	for index, number := range numbers {
		fmt.Printf("Student Name: %v, Marks : %v \n", index, number)
	}

}
