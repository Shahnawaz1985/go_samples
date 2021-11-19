package main

import (
	"fmt"
	"time"
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

func reportNap(name string, delay int) {
	for i := 0; i < delay; i++ {
		fmt.Println(name, "sleeping")
		time.Sleep(1 * time.Second)
	}
	fmt.Println(name, "wakes up!")
}

func sendValue(myChannel chan string) {
	reportNap("sending go routine", 2)
	fmt.Println("***sending value***")
	myChannel <- "a"
	fmt.Println("***sending value***")
	myChannel <- "ab"
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
	fmt.Println()
	fmt.Println("--------------------")

	testChannel := make(chan string)
	go sendValue(testChannel)
	reportNap("receiving goroutine", 5)
	fmt.Println(<-testChannel)
	fmt.Println(<-testChannel)
}
