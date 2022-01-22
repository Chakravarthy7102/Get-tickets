package main

import (
	"fmt"
	"strings"
)

func main() {
	companyName := "TicketsGet"
	const totalTickets = 50
	var availableTickets int = 50
	bookings := []string{}

	fmt.Printf("hello world! \nWelcome To %s \n", companyName)
	fmt.Printf("Total Tickets are %d Available Tickets %d \n", totalTickets, availableTickets)
	fmt.Println("Book yo Tickets fast..very fasst")

	for {
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

		for availableTickets-ticketsBooked < 0 {
			fmt.Printf("Please Enter the ticket in range of %d : ", availableTickets)
			fmt.Scan(&ticketsBooked)
		}
		availableTickets = availableTickets - ticketsBooked

		// bookings[0] = firstName + " " + lastName
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("You %s %s have booked %d tickets you will get a confirmation at this %s mail \n", firstName, lastName, ticketsBooked, email)

		fmt.Printf("%d remaining ...\n", availableTickets)

		firstNames := []string{}

		for _, booking := range bookings {

			firstname := strings.Fields(booking)
			firstNames = append(firstNames, firstname[0])

		}
		fmt.Println(firstNames)

		if availableTickets <= 0 {
			fmt.Println("SOLD OUT!!")
			break
		}

		fmt.Printf("All users %v \n", bookings)
	}
}
