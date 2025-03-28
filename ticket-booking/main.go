package main

import (
	"fmt"
	"strings"
)

func book_ticket() {
	fmt.Println("hey this is Go Lang hello program.")
	var firstName string
	var lastName string
	var email string
	var totalTickets uint = 100
	var remainingTicket uint
	var tickets uint

	fmt.Print("Enter first name::")
	fmt.Scan(&firstName)
	fmt.Print("Enter last name::")
	fmt.Scan(&lastName)
	fmt.Print("Enter email::")
	fmt.Scan(&email)
	fmt.Print("Number of ticket::")
	fmt.Scan(&tickets)
	if tickets <= totalTickets && tickets > 0 {
		remainingTicket = totalTickets - tickets
		fmt.Printf("Hi %v %v, your %v tickets have been booked and will be send to %v soon\n Happy booking..\n", firstName, lastName, tickets, email)
		fmt.Printf("Available tickets are :: %v\n", remainingTicket)
	} else {
		remainingTicket = totalTickets
		fmt.Println("You have requested for booking more than remaining tickets")
		fmt.Printf("Remaining tickets are :: %v\n", remainingTicket)
	}
}

func practice_slice() {
	var numbers []int
	var count int = 5

	for i := 0; i < count; i++ {
		var temp int
		fmt.Printf("Enter %v element...", i+1)
		fmt.Scan(&temp)
		numbers = append(numbers, temp)
	}

	fmt.Print("Entire array::", numbers)

	fmt.Print("Fist element::", numbers[0])
}

func practice_array() {
	var numbers [10]int
	var count int = 10

	for i := 0; i < count; i++ {
		fmt.Printf("Enter %v element...", i+1)
		fmt.Scan(&numbers[i])
	}

	fmt.Print("Entire array::", numbers)

	fmt.Print("Fist element::", numbers[0])

}

func string_function() {
	fullNames := []string{}
	count := 5
	for i := 0; i < count; i++ {
		var fistName string
		var lastName string
		fmt.Printf("Enter %v first name::\n", i+1)
		fmt.Scan(&fistName)
		fmt.Printf("Enter %v last name::\n", i+1)
		fmt.Scan(&lastName)
		fullNames = append(fullNames, fistName+" "+lastName)
	}

	firstNames := []string{}

	for _, fullName := range fullNames {
		names := strings.Fields(fullName)
		firstNames = append(firstNames, names[0])
	}

	fmt.Print("All first Names are::", firstNames, "\n")

}

func control_statement() {
	availableTicket := 0
	fmt.Print("Enter the num of availableTicket::")
	fmt.Scan(&availableTicket)
	for {
		if availableTicket <= 0 {
			break
		}
		var temp int
		fmt.Print("Enter the ticket::")
		fmt.Scan(&temp)
		if temp < 1 {
			fmt.Println("Invalid ticket entered !")
			continue
		}
		if availableTicket >= temp {
			availableTicket = availableTicket - temp
			fmt.Println("Ticket Boocked !")

		} else {
			fmt.Println("Ooops...! you booked more tickets than available !")

		}

	}
	fmt.Println("Tickets Exhausted !\nTry next time ")
}

func switch_case() {
	var invalidCity bool = false
	for {
		var city string
		fmt.Print("Enter destiation ticket::")
		fmt.Scan(&city)

		switch city {
		case "Kolkata", "Hawrah":
			fmt.Println("booking ticket for kolkata or Hawrah")

		case "ballia", "siwan":
			fmt.Println("Booking ticket for Ballia or Siwan")

		default:
			fmt.Println("Booking not started for ", city)
			invalidCity = true
		}
		if invalidCity {
			break
		}
	}

}

func main() {
	// Variables
	book_ticket()

	// Array (fixed size)
	practice_array()

	// Slice (dynamic array)
	practice_slice()

	// string

	string_function()

	// control Statements
	control_statement()

	// switch statement
	switch_case()
}
