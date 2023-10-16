package main

import (
	"fmt"

	// "rsc.io/quote"
	"gogo/greetings"
)

// func Hello(name string) string {
// 	message := fmt.Sprintf("Hi %v. Welcome!", name)
// 	return message
// }

func main() {

	message := greetings.Hello("Aileen")
	fmt.Println(message)
}