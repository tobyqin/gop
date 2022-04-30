package main

import (
	"fmt"
	"log"

	"tobyqin.cn/greetings"
)

func main() {
	log.SetPrefix("grettings: ")
	log.SetFlags(0)

	for i := 0; i < 5; i++ {
		message, err := greetings.Hello("Toby")

		if err != nil {
			log.Fatal(err)
		}

    fmt.Println(message)

	}
}
