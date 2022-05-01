package main

import (
	"fmt"
	"strings"
)

func main() {
	var appName = "booking app"
	const totalTickets uint = 100
	var remainingTickets uint = totalTickets
	var users []string

	fmt.Printf("type of appName: %T\n", appName)
	println("Address of appName", &appName)

	fmt.Printf("Welcome to %v!\n", appName)
	fmt.Printf("There are %v tickets remaining.\n", remainingTickets)

	// var userName string
	// var userTickets int

	// fmt.Printf("Please enter your name: ")
	// fmt.Scan(&userName)
	// fmt.Printf("Hello %v, how many tickets would you like? ", userName)
	// fmt.Scan(&userTickets)
	// fmt.Printf("\n%v, you have %v tickets.", userName, userTickets)

	var bookings = [50]string{"nana q", "nicole f", "petter g"}
	fmt.Printf("\nfull names%v\n", bookings)
	users = append(users, "abc")
	firstNames := []string{}

	for _, booking := range bookings {
		if booking != "" {
			names := strings.Split(booking, " ")
			firstNames = append(firstNames, names[0])
		}

	}

	fmt.Printf("firstNames: %v\n", firstNames)

	fmt.Printf("\n%v\n", users)
	// fmt.Printf("\n%T", users)
	// fmt.Printf("\n%v", len(users))

	isValidName := len(appName) > 0
	isValidEmail := strings.Contains(appName, "@")
	fmt.Printf("\n%v, %v\n", isValidName, isValidEmail)

	var c rune = '中'
	fmt.Printf("%v %c", c, c)

}
