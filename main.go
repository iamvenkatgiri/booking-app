package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v Booking Application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend the conference")

	var bookings []string
	for {
		if remainingTickets > 0 {
			var firstName string
			var lastName string
			var userTickets uint
			var email string

			fmt.Printf("Enter the number of tickets(Available Tickets:%v)", remainingTickets)
			fmt.Scanf("%d", &userTickets)

			if userTickets <= remainingTickets {
				fmt.Println("Enter your first name")
				fmt.Scan(&firstName)
				fmt.Println("Enter your last name")
				fmt.Scan(&lastName)
				fmt.Println("Enter the email")
				fmt.Scan(&email)
				remainingTickets = remainingTickets - userTickets
				bookings = append(bookings, firstName+" "+lastName)

				fmt.Printf("Booking Confirmed for %v %v.\n%v Tickets Booked. Confirmation email will be sent to %v \n", firstName, lastName, userTickets, email)
				fmt.Printf("%v Remaining tickets for %v\n", remainingTickets, conferenceName)

				firstNames := []string{}

				for _, booking := range bookings {

					var names = strings.Fields(booking)

					firstNames = append(firstNames, names[0])
				}
				fmt.Printf("The first Names of bookings are : %v \n", firstNames)

			} else {
				fmt.Printf("We only have %v tickets. You can't book %v tickets \n", remainingTickets, userTickets)
				fmt.Println("Try Again")
				continue
			}

		} else {
			fmt.Printf("All the tickets are booked. Remaining tickets = %v \n", remainingTickets)
			break
		}

	}

}
