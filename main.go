package main

import (
	"fmt"
	"time"
)

const conferenceTickets int = 50
var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]userData , 0) 

type userData struct {
	firstName string
	lastName  string
	email 	  string
	numberOfTickets uint
}

func main() {
  	 greetUsers()

	for remainingTickets > 0 && len(bookings) < 50  {  
	    
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isEmailValid , isValidTickets := ValidateUserInputs(firstName, lastName, email, userTickets, remainingTickets)
	
		if  isValidName && isEmailValid && isValidTickets {
		     			
		    bookTicket(userTickets , firstName , lastName , email)

			go  sendTicket(userTickets , firstName , lastName , email)
	
	        firstNames := getFirstNames()  
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
  
func greetUsers( ) {
	fmt.Printf("Welcome to %v booking app\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets , remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{} 

	for _, booking := range bookings {
		firstNames =  append(firstNames, booking.firstName )
	}

	return firstNames 		
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

func bookTicket(userTickets uint, firstName string, lastName string,  email string) {
	remainingTickets -=  userTickets 

    var userData  = userData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets, 
	}

	bookings = append(bookings, userData)
	fmt.Printf("userData :  %v", userData)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will get an email to %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n.", remainingTickets, conferenceName) 
} 


func sendTicket(userTickets uint, firstName string, lastName string , email string)  {
  time.Sleep(10 * time.Second)	
  var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
  fmt.Println("###########")
  fmt.Printf("Sending ticket:\n %v \n to email address %v\n", ticket, email)
  fmt.Println("###########")
}