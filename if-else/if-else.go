package main

import "fmt"

func main() {
	a := 0

	if a > 5 {
		fmt.Printf("%v is grater than 5 \n", a)
	} else if a < 5 && a > 1 {
		fmt.Printf("%v is smaller than 5 \n", a)
	} else {
		fmt.Println("A is number", a)
	}
}
