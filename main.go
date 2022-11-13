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

	fmt.Printf("Welcome to %v booking app\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets , remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for remainingTickets > 0 && len(bookings) < 50  {  
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
	
		if(userTickets <= remainingTickets) {
		
			remainingTickets -=  userTickets 

			bookings = append(bookings, firstName + " " + lastName)
		
			fmt.Printf("Thank you %v %v for booking %v tickets. You will get an email to %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n.", remainingTickets, conferenceName)
	
	
			firstNames := []string{}
	
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
	
			fmt.Printf("The first names of booking ares : %v\n ", firstNames)
	
			if remainingTickets == 0 {
				  fmt.Println("Our conference booked out.Come back next year")
				  break
			}	

		}  else {
			fmt.Printf("We only have %v ticket remaining, so you cant book %v tickets\n", remainingTickets, userTickets)  
		}	
	}	
} 
