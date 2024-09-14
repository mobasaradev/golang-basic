package main

import "fmt"

func main() {
	for i := 0; i <= 3; i++ {
		fmt.Println(i)
	}

	m := 0
	for {
		fmt.Println("loop")
		m++
		if m == 5 {
			break
		}
	}

	nums := []int{50, 49, 48, 47, 46, 45}
	for index, num := range nums {
		fmt.Printf("Index %d : Value %d \n", index, num)
	}

	names := []string{"Hello", "Sara", "you", "want", "to", "be", "Evan"}
	for index, char := range names {
		fmt.Printf("Index : %d, Value : %s\n", index, char)
	}
}
