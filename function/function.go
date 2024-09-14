package main

import "fmt"

func main() {
	printString()
	sumResult := sum(4,1)
	fmt.Printf("The result of the sum : %v \n", sumResult)

	multiResult := multi(3,4)
	fmt.Printf("The result of the multi : %v \n", multiResult)
}

func sum(a,b int) int {
	return a + b
}

func multi(a,b int) (result int) {
	result = a * b
	return 
}

func printString()  {
	fmt.Println("This is function test")
}