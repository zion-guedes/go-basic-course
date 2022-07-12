package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

//Package level vars
var conferenceName = "Go Conference 2022"
var remainingTickets uint = 50

//var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)

const conferenceTickets = 50

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
	isNews          bool
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	//for {

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		//Print first names
		fmt.Printf("These are all our bookings: %v\n", getFirstNames())

		noTicketsRemaining := remainingTickets == 0
		if noTicketsRemaining {
			//end the loop
			fmt.Println("Our Conference is booked out. Until next year!")
			//break
		}
	} else {
		//User input validation
		if !isValidName {
			fmt.Println("Your first or last name is too short. Try again!")
		}
		if !isValidEmail {
			fmt.Println("Your email is not valid. Try again!")
		}
		if !isValidTicketNumber {
			fmt.Println("Your number of tickets is invalid. Try again!")
		}
	}
	wg.Wait()
	//}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here")
	fmt.Println("---------------------")
}

func getFirstNames() []string {
	//Get only first names to show
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	//Ask user for input
	fmt.Println("Please enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email name: ")
	fmt.Scan(&email)

	fmt.Println("Please enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	//Create a map for a user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("These are all our bookings: %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v \n", userTickets, firstName, lastName)
	fmt.Println("#############")
	fmt.Printf("Sending ticket %v to email %v\n", ticket, email)
	fmt.Println("#############")
	wg.Done()
}
