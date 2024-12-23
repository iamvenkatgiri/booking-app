package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var RemainingTickets uint = 50
var bookings = make([]UserData, 0)

// customized data type to which can contain diff data types unlike map
type UserData struct {
	firstName   string
	lastName    string
	email       string
	noOfTickets uint
}

var wg = sync.WaitGroup{}

// Main Function
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
				wg.Add(1)
				go sendTickets(userTickets, firstName, lastName, email)

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
	wg.Wait()
}

// Function to Greet Users
func greetUsers() {
	fmt.Printf("Welcome to %v Booking App\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, RemainingTickets)
	fmt.Println("Get your tickets here to attend the conference")
}

// Function to extract firstNames
func getFirstNames(bookings []UserData) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

// Function to get user data
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

// Function to book tickets
func bookTickets(userTickets uint, firstName string, lastName string, email string) (uint, []UserData) {

	RemainingTickets = RemainingTickets - userTickets

	var userData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		noOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Println("Bookings List:", bookings)
	fmt.Printf("Booking Confirmed for %v %v.\n%v Tickets Booked. Confirmation email will be sent to %v \n", firstName, lastName, userTickets, email)

	return RemainingTickets, bookings
}

func sendTickets(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("######################")
	fmt.Printf("Sending Tickets:\n%v \n to email address %v\n", ticket, email)
	fmt.Println("######################")
	wg.Done()
}
