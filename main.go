package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	 bookings := []string{}

	 greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for remainingTickets > 0 && len(bookings) < 50  {  
	    
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isEmailValid , isValidTickets := validateUserInputs(firstName, lastName, email, userTickets, remainingTickets)
	
		if  isValidName && isEmailValid && isValidTickets {
		     			
		    bookTicket(remainingTickets , userTickets , firstName , lastName , conferenceName , email , bookings )
	
	        firstNames := getFirstNames(bookings)  
			fmt.Printf("The first names of booking ares : %v\n ", firstNames)
			
			if remainingTickets == 0 {
				  fmt.Println("Our conference booked out.Come back next year")
				  break
			}	

		}  else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short.")  
			}
			if !isEmailValid {
				fmt.Println("email is invalid.")  
			}
			if !isValidTickets {
				fmt.Println("User tickets you entered is invalid.")
			}		
		}	
	}	
} 
 
func greetUsers(conferenceName string,conferenceTickets int ,remainingTickets uint) {
	fmt.Printf("Welcome to %v booking app\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets , remainingTickets)
	fmt.Println("Get your tickets here to attend")
}


func getFirstNames(bookings []string) []string {
	firstNames := []string{}

	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	return firstNames 		
}


func validateUserInputs(firstName string, lastName string, email string, userTickets  uint , remainingTickets uint) (bool, bool, bool)  { 
	isValidName := len(firstName) >= 2 && len(lastName) >= 2  
	
	isEmailValid := strings.Contains(email, "@")

	isValidTickets := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isEmailValid , isValidTickets
}


func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string 
	var userTickets uint
	//ask user for their name
	fmt.Println("Enter your first name:  ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:  ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:  ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:  ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(remainingTickets uint, userTickets uint, firstName string, lastName string, conferenceName string, email string, bookings []string) {
	remainingTickets -=  userTickets 

	bookings = append(bookings, firstName + " " + lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will get an email to %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n.", remainingTickets, conferenceName) 
} 