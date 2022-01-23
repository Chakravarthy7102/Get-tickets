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

		//user validation

		if !((len(firstName) >= 3) && (len(lastName) >= 3)) {
			fmt.Println("Please Enter a Valid Name :)")
			fmt.Println("##########################################################")
			continue
		}

		//email

		check := strings.Contains(email, "@")
		check1 := strings.Contains(email, ".com")

		if !(check && check1) {
			fmt.Println("Please enter a valid fomat of your email")
			fmt.Println("##########################################################")
			continue
		}

		//tickets

		if !(ticketsBooked > 0) {
			fmt.Println("Please enter a valid amount of Tickets")
			fmt.Println("##########################################################")
			continue
		}

		for availableTickets-ticketsBooked < 0 {
			fmt.Printf("Please Enter the ticket in range of %d : \n", availableTickets)
			fmt.Println("##########################################################")
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
		fmt.Println("##########################################################")
		fmt.Println("##########################################################")
	}
}
