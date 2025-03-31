package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

var WG = sync.WaitGroup{}

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

	// function with parameter
	greetUser("Ramesh")

	// function with return value
	num := 10.1
	fmt.Printf("square of %v is %v\n", num, square(num))

	// function with multiple returns
	greet, strLen := say_hi("vivek")
	fmt.Printf("%v\nstring  Len::%v\n", greet, strLen)

	// use of map

	fmt.Println(use_map("vivek", "verma", 20))

	// Multi threding
	WG.Add(1)
	go greetUser("vivek kumar verma")
	WG.Add(1)
	go print_square(10.1)

	WG.Add(1)
	go greetUser("sarma verma")
	WG.Add(1)
	go print_square(20.1)
	WG.Wait()

}

func greetUser(userName string) {
	time.Sleep(5 * time.Second)
	fmt.Printf("Hi %v,\nWelcome to ticket booking portal\n", userName)
	WG.Done()
}

// function with return values

func print_square(l float64) {
	// print the area of The square
	time.Sleep(10 * time.Second)
	fmt.Printf("Area =%v\n", l*l)
	WG.Done()
}

func square(l float64) float64 {
	// returns the area of The square

	return l * l
}

// function with Multiple returns

func say_hi(name string) (string, int) {
	var greeting string = "Hi " + name + ",\nHope you are having a Good day !"
	return greeting, len(greeting)

}

// function to demonstrate the use of Map (dictionary in python).
// Declare map using make for empty map declaration
// go map[key_datatype]value_datatype
func use_map(firstName string, lastName string, age int64) map[string]string {
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["age"] = strconv.FormatUint(uint64(age), 10)
	return userData
}
