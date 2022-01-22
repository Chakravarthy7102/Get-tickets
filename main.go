package main

import "fmt"

func main() {
	companyName := "TicketsGet"
	const totalTickets = 50
	var availableTickets int = 50
	fmt.Printf("hello world! \nWelcome To %s \n", companyName)
	fmt.Printf("Total Tickets are %d Available Tickets %d \n", totalTickets, availableTickets)
	fmt.Println("Book yo Tickets fast..very fasst")

	var firstName string
	var lastName string
	var email string
	var ticketsBooked int

	fmt.Println("Enter Your first name")
	fmt.Scan(&firstName)
	fmt.Println("Enter Your last name")
	fmt.Scan(&lastName)
	fmt.Println("Enter your Email")
	fmt.Scan(&email)
	fmt.Println("Enter no of ticket")
	fmt.Scan(&ticketsBooked)

	availableTickets = availableTickets - ticketsBooked

	fmt.Printf("You %s %s have booked %d tickets you will get a confirmation at this %s mail \n", firstName, lastName, ticketsBooked, email)

	fmt.Printf("%d remaining ...\n", availableTickets)
}
