package userValidation

import "strings"

// valid the user input
func ValidUserInput(firstName, lastName, email string, remainingTickets, userTickets int) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidSeatsNum := userTickets > 0 && remainingTickets >= userTickets
	return isValidEmail, isValidName, isValidSeatsNum
}
