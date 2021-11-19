package main

import (
	"fmt"
)

func greetChannel(mychannel chan string) {
	mychannel <- "hi"
}

func printAbc(channel chan string) {
	channel <- "a"
	channel <- "b"
	channel <- "c"
}

func printDef(channel chan string) {
	channel <- "d"
	channel <- "e"
	channel <- "f"
}

func main() {
	fmt.Println("Performing channel test.")
	mychan := make(chan string)
	go greetChannel(mychan)
	tempVar := <-mychan
	fmt.Println("value received from channel :", tempVar)

	channel11 := make(chan string)
	channel12 := make(chan string)

	go printAbc(channel11)
	go printDef(channel12)

	fmt.Print(<-channel11)
	fmt.Print(<-channel12)
	fmt.Print(<-channel11)
	fmt.Print(<-channel12)
	fmt.Print(<-channel11)
	fmt.Print(<-channel12)
}
