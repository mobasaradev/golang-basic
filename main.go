package main

import (
	"fmt"
	"strings"
)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets int
}

var orgName string = "GoLang Org"
var bookings = make([]UserData, 0)

const totalTickets int = 50

var remainingTickets int = 50

func main() {
	greeting()

	for {
		// User input
		firstName, lastName, email, userTickets := getUserInput()
		// user input validation
		isValidName, isValidEmail, isValidSeatsNum := validUserInput(firstName, lastName, email, remainingTickets, userTickets)

		if isValidName && isValidEmail && isValidSeatsNum {
			// booked tickets
			bookedTicket(firstName, lastName, email, userTickets)
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

func bookedTicket(firstName, lastName, email string, userTickets int) {
	// check remaining tickets
	remainingTickets = remainingTickets - userTickets
	userData := UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	// list of user - append func for add data in a array for array or slices.
	bookings = append(bookings, userData)
	// congratulation msg to user
	fmt.Printf("Thank you %v for collecting %v tickets. we'll contact you in %v.\n", firstName, userTickets, email)
	fmt.Printf("list of users name : %v\nRemaining tickets number : %v\n", bookings, remainingTickets)
}

// valid the user input
func validUserInput(firstName, lastName, email string, remainingTickets, userTickets int) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidSeatsNum := userTickets > 0 && remainingTickets >= userTickets
	return isValidEmail, isValidName, isValidSeatsNum
}
