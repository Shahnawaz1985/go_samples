package main

import (
	"mutextest"
	"time"
)

func main() {
	r := mutextest.NewRoverDriver()
	time.Sleep(3 * time.Second)
	r.Left()
	time.Sleep(3 * time.Second)
	r.Right()
	time.Sleep(3 * time.Second)
}
