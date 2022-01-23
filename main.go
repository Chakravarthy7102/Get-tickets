package main

import (
	"fmt"
	"get-tickets/subpack"
	"strings"
	"sync"
	"time"
)

var wl = sync.WaitGroup{}

func main() {
	companyName := "TicketsGet"
	const totalTickets = 50
	var availableTickets int = 50
	type UserData struct {
		firstName     string
		lastName      string
		email         string
		ticketsBooked int
	}
	bookings := make([]UserData, 0)

	//struct

	subpack.UserGreetings(companyName, totalTickets, availableTickets)

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
			subpack.EndLine()
			continue
		}

		//email

		check := strings.Contains(email, "@")
		check1 := strings.Contains(email, ".com")

		if !(check && check1) {
			fmt.Println("Please enter a valid fomat of your email")
			subpack.EndLine()
			continue
		}

		//tickets

		if !(ticketsBooked > 0) {
			fmt.Println("Please enter a valid amount of Tickets")
			subpack.EndLine()
			continue
		}

		for availableTickets-ticketsBooked < 0 {
			fmt.Printf("Please Enter the ticket in range of %d : \n", availableTickets)
			subpack.EndLine()
			fmt.Scan(&ticketsBooked)
		}

		availableTickets = availableTickets - ticketsBooked

		userData := UserData{
			firstName:     firstName,
			lastName:      lastName,
			email:         email,
			ticketsBooked: ticketsBooked,
		}

		//maps

		// bookingMap := make(map[string]string)

		// bookingMap["first Name"] = firstName
		// bookingMap["Last Name"] = lastName
		// bookingMap["email"] = email
		// bookingMap["Number of Tickets"] = strconv.FormatInt(int64(ticketsBooked), 10)

		// bookings[0] = firstName + " " + lastName
		bookings = append(bookings, userData)

		fmt.Printf("You %s %s have booked %d tickets you will get a confirmation at this %s mail \n", firstName, lastName, ticketsBooked, email)

		fmt.Printf("%d remaining ...\n", availableTickets)

		wl.Add(1)
		go TicketToMail(ticketsBooked, firstName, lastName, email)

		//slices
		firstNames := []string{}

		for _, booking := range bookings {

			firstname := booking.firstName
			firstNames = append(firstNames, firstname)

		}
		fmt.Println(firstNames)

		if availableTickets <= 0 {
			fmt.Println("SOLD OUT!!")
			break
		}

		fmt.Printf("All users %v \n", bookings)
		subpack.EndLine()
		wl.Wait()

	}
}

func TicketToMail(ticketsBooked int, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	fmt.Println("##########################################")
	message := fmt.Sprintf("%v tickets are booked by %v %v to mail %v \n", ticketsBooked, firstName, lastName, email)
	fmt.Printf("Mailed Succesfully!! %v\n", message)
	fmt.Println("##########################################")

	wl.Done()

}
