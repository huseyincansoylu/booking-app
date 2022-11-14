package main

import "strings"

func ValidateUserInputs(firstName string, lastName string, email string, userTickets  uint , remainingTickets uint) (bool, bool, bool)  { 
	isValidName := len(firstName) >= 2 && len(lastName) >= 2  
	
	isEmailValid := strings .Contains(email, "@")

	isValidTickets := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isEmailValid , isValidTickets
}