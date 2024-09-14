package main

import "fmt"

func main() {
	num := 3
	switch num {
	case 1:
		fmt.Println("Sunday")
	case 2:
		fmt.Println("Monday")
	case 3:
		fmt.Println("Tuesday")
	default:
		fmt.Println("Weekend")
	}

	month := "January"

	switch month {
	case "January", "February", "March":
		fmt.Println("Winter Season")
	case "April", "May":
		fmt.Println("Summer Season")
	default:
		fmt.Println("Unknown season")
	}

	temperature := 30

	switch {
	case temperature <= 0:
		fmt.Println("Freezing")
	case temperature > 1 && temperature <= 18:
		fmt.Println("Cold")
	case temperature >= 19 && temperature <= 25:
		fmt.Println("warm")
	case temperature >= 26:
		fmt.Println("Hot")
	default:
		fmt.Println("death")
	}

}
