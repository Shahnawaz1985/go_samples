package main

import (
	"fmt"
	"greeting"
	"keyboard"
	"log"
)

func main() {
	fmt.Println("Test Hello!")
	greeting.Greet1()
	greeting.Greet2()

	fmt.Println("Enter a grade: ")

	grade, err := keyboard.GetFloat()

	if err != nil {
		log.Fatal(err)
	}

	var status string

	if grade > 60 {
		status = "passing score"
	}else{
		status = "failing score"
	}

	fmt.Println("A grade of", grade, "is", status)
}