package main

import (
	// "bufio"
	"fmt"
	// "os"
	// "strings"
)

func main() {
	var orgName string = "GoLang"
	const totalSeats int = 50
	var remainSeats uint = 50
	var fullName string
	var bookedSeat uint
	// List - where we need to mention the length
	// var bookingSeats [5]string

	// Slices - length will be dynamic
	// bookingSeats := []string{}
	var bookingSeats []string

	fmt.Printf("Welcome to %v, We've total %v seats.\n", orgName, totalSeats)

	for {

		fmt.Println("*Enter your name : ")
		fmt.Scan(&fullName)
		// reader := bufio.NewReader(os.Stdin)
		// fullName, _ := reader.ReadString('\n')

		fmt.Println("How many seat do you want to book? ")
		fmt.Scan(&bookedSeat)

		// isValidEmail := strings.Contains(email, "@")
		isValidName := len(fullName) >= 2
		isValidSeatsNum := bookedSeat > 0 && bookedSeat <= remainSeats

		if isValidName && isValidSeatsNum {
			remainSeats = remainSeats - bookedSeat
			// List - need to mention the index
			// bookingSeats[0] = fullName

			// Slices - no need to mention the index
			bookingSeats = append(bookingSeats, fullName)

			fmt.Printf("Thank you %v for booking %v tickets\n", fullName, bookedSeat)
			fmt.Printf("%v tickets are left for %v conference \n", remainSeats, orgName)
			fmt.Printf("The Students whom booked the seats are %v.\n", bookingSeats)

			if remainSeats == 0 {
				fmt.Println("We're sold out.")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("You've enter invalid name")
			}
			if !isValidSeatsNum {
				fmt.Println("You've enter invalid seat booking number")
			}

		}
	}
}
