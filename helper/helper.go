package helper

import (
	"strings"
)

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, RemainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
	isValidInput := userTickets > 0 && userTickets <= RemainingTickets
	return isValidEmail, isValidInput, isValidName
}
