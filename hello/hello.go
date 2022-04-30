package main

import (
	"fmt"
  "log"

	"tobyqin.cn/greetings"
)

func main (){
  log.SetPrefix("grettings: ")
  log.SetFlags(0)

  message, err := greetings.Hello("")

  if err !=nil{
    log.Fatal(err)
  }

  fmt.Println(message)
}
