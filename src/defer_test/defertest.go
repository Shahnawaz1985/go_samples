package main

import (
	"fmt"
)

func Socialize() {
	defer fmt.Println("Hello There!")
	fmt.Println("Hello")
	fmt.Println("Hello Again")
}

func main() {
	Socialize()
}
