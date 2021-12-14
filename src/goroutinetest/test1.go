package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c)
	}

	timeout := time.After(4 * time.Second)
	for i := 0; i < 5; i++ {
		select {
		case gopherID := <-c:
			fmt.Println("gopher", gopherID, " has finished sleeping")
		case <-timeout:
			fmt.Println("timing out!")
			return
		}
	}
	time.Sleep(4 * time.Second)
}

func sleepyGopher(id int, c chan int) {
	//time.Sleep(3 * time.Second)
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	fmt.Println("***", id, " snore***")
	c <- id
}
