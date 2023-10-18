package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50
var bookingList = make([]UserData, 0)

type UserData struct {
	firstName   string
	lastName    string
	email       string
	noOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUser()

	// Infinite Loop

	firstName, lastName, email, userTickets := getUserInputs()

	isValidName, isValidEmail, isValidUserTickets := helper.ValidateInputs(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidUserTickets {

		bookTickets(userTickets, firstName, lastName, email)

		wg.Add(1) // Set the no of go routines to wait
		go mailTickets(firstName, lastName, email, userTickets)

		firstNames := getFirstName()
		fmt.Printf("First Names are %v \n", firstNames)

		if remainingTickets == 0 {
			fmt.Printf("Our conference is booked out. Come back next year.\n")
		}
	} else {

		if !isValidName {
			fmt.Printf("Please enter a valid first name or last name\n")
		} else if !isValidEmail {
			fmt.Printf("Please enter a valid email\n")
		} else if !isValidUserTickets {
			fmt.Printf("Please enter a valid user tickets\n")
		}

		fmt.Printf("Your input is invalid please try again\n")
	}

	wg.Wait() // Blocks until all the go routines are finished

}

func greetUser() {
	fmt.Printf("Welcome to %v booking application! \n", conferenceName)
	fmt.Printf("We have total of %v tickets and  %v remaining tickets \n", conferenceTickets, remainingTickets)
	fmt.Println("Get you your tickets here and attend the conference")
}

func getFirstName() []string {
	// Get the firstNames List
	firstNames := []string{}

	// range allow use to iterate over different data structures
	// range for array and slice give use index as well as each element value
	for _, booking := range bookingList {
		firstNames = append(firstNames, booking.firstName) // Append the first name to the firstName list append(slice,element)
	}

	return firstNames
}

func getUserInputs() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first Name")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last Name")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address")
	fmt.Scan(&email)
	fmt.Println("Enter your number of tickets")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// Create a new map for the user
	var userMap = UserData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		noOfTickets: userTickets,
	}

	bookingList = append(bookingList, userMap)

	fmt.Printf("List of bookings: %v\n", bookingList)
	fmt.Printf("Thanks you %v %v for booking %v tickets. You will receive confirmation on your email %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("Remaining tickets availible are %v for %v\n", remainingTickets, conferenceName)
}

func mailTickets(firstName string, lastName string, email string, userTickets uint) {

	var userTicket = fmt.Sprintf("%v %v have booked %v tickets", firstName, lastName, userTickets)
	time.Sleep(10 * time.Second)
	fmt.Println("################################################################")
	fmt.Printf("Thank you %v for %v and it has been send to your email %v\n", userTicket, conferenceName, email)
	fmt.Println("################################################################")

	wg.Done() // Decrement the wait group counter

}
