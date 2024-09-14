package main

import "fmt"

func main() {
	var numbers = [5]int{}
	// var numbers [5]int
	numbers[2] = 2
	fmt.Println("Numbers : ", numbers)
	fmt.Println("Numbers : ", numbers[1])
	fmt.Println("Numbers : ", len(numbers))
}
