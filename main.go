package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var RemainingTickets uint = 50
var bookings = make([]map[string]string, 0)

func main() {

	greetUsers()

	for {
		if RemainingTickets > 0 {

			firstName, lastName, email, userTickets := getUserData()
			isValidEmail, isValidInput, isValidName := helper.ValidateUserInput(firstName, lastName, email, userTickets, RemainingTickets)

			if isValidEmail && isValidInput && isValidName {

				RemainingTickets, bookings := bookTickets(userTickets, firstName, lastName, email)
				fmt.Printf("%v Remaining tickets for %v\n", RemainingTickets, conferenceName)
				firstNames := getFirstNames(bookings)
				fmt.Printf("The first Names of bookings are : %v \n", firstNames)

			} else {
				if !isValidEmail {
					fmt.Printf("Email Entered: %v\nMake sure the email contains '@' and '.' in it.\n", email)
				}
				if !isValidName {
					fmt.Printf("First Name Entered: %v\nLast Name Entered:%v\nLength should be minimum two characters\n", firstName, lastName)
				}
				if !isValidInput {
					fmt.Printf("The requested tickets:%v are more than the available tickets:%v\n", userTickets, RemainingTickets)
				}
			}

		} else {
			fmt.Printf("All the tickets are booked. Remaining tickets = %v \n", RemainingTickets)
			break
		}

	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v Booking App\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, RemainingTickets)
	fmt.Println("Get your tickets here to attend the conference")
}

func getFirstNames(bookings []map[string]string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
}

func getUserData() (string, string, string, uint) {
	var firstName string
	var lastName string
	var userTickets uint
	var email string

	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter the email")
	fmt.Scan(&email)

	fmt.Printf("Enter the number of tickets(Available Tickets:%v)", RemainingTickets)
	fmt.Scanf("%d", &userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) (uint, []map[string]string) {

	RemainingTickets = RemainingTickets - userTickets

	// Create a map for a user
	var userData = make(map[string]string)

	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["noOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Println("Bookings List:", bookings)
	fmt.Printf("Booking Confirmed for %v %v.\n%v Tickets Booked. Confirmation email will be sent to %v \n", firstName, lastName, userTickets, email)

	return RemainingTickets, bookings
}
