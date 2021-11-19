package main

import (
	"fmt"
	"greeter"
	"keyboard"
	"log"
)

func main() {

	fmt.Println("Hello World!")
	greeter.Greet1()
	greeter.Greet2()
	fmt.Println("")
	userInput, err := keyboard.GetInput()

	if err != nil {
		log.Fatal(err)
	}

	var status string

	if userInput > 60 {
		status = "passing"
	} else {
		status = "failing"
	}

	fmt.Println("\nA grade of", userInput, "is", status, "grade!")

}
