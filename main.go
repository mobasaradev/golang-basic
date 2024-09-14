package main

import (
	"fmt"
	userValidation "go-sara/user-validation"
)

var orgName string = "GoLang Org"
var users []string

const totalTickets int = 50

var remainingTickets int = 50

func main() {
	greeting()

	for {
		// User input
		firstName, lastName, email, userTickets := getUserInput()
		// user input validation
		isValidName, isValidEmail, isValidSeatsNum := userValidation.ValidUserInput(firstName, lastName, email, remainingTickets, userTickets)

		if isValidName && isValidEmail && isValidSeatsNum {
			// booked tickets
			bookedTicket(firstName, email, userTickets)
			// stop selling tickets if remaining ticket is 0
			if remainingTickets == 0 {
				fmt.Println("We're Sold Out")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("You've enter invalid name")
			}
			if !isValidEmail {
				fmt.Println("You've enter invalid email")
			}
			if !isValidSeatsNum {
				fmt.Println("You've enter 0 ticket. Enter a valid number of tickets.")
			}
		}
	}
}

func greeting() {
	fmt.Println("Welcome to", orgName)
	fmt.Printf("We've %v tickets. Available tickets : %v\n", totalTickets, remainingTickets)
}



func getUserInput() (string, string, string, int) {
	var firstName, lastName, email string
	var userTickets int

	// user input
	fmt.Println("Enter your first name :")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name :")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email name :")
	fmt.Scan(&email)
	fmt.Println("Number of tickets you want take ?")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookedTicket(firstName, email string, userTickets int,) {
	// list of user - append func for add data in a array for array or slices.
	users = append(users, firstName)
	// check remaining tickets
	remainingTickets = remainingTickets - userTickets

	// congratulation msg to user
	fmt.Printf("Thank you %v for collecting %v tickets. we'll contact you in %v.\n", firstName, userTickets, email)
	fmt.Printf("list of users name : %v\nRemaining tickets number : %v\n", users, remainingTickets)
}
