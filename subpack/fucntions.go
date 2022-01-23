package subpack

import (
	"fmt"
)

func UserGreetings(companyName string, totalTickets int, availableTickets int) {

	fmt.Printf("hello world! \nWelcome To %s \n", companyName)
	fmt.Printf("Total Tickets are %d Available Tickets %d \n", totalTickets, availableTickets)
	fmt.Println("Book yo Tickets fast..very fasst")

}

func EndLine() {
	fmt.Println("##########################################################")
	fmt.Println("##########################################################")
}
