package main

import (
	"fmt"
	"log"

	"gop-app/booking"
	"gop-app/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	for i := 0; i < 5; i++ {
		message, err := greetings.Hello("Toby")

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(message)
	}
	booking.Book()
	booking.CreateUserDataStruct()
}
