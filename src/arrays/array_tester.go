package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Array initialization :")

	var notes [6]string

	fmt.Println("Empty array", notes)

	notes[0] = "Hello"
	notes[1] = "World"
	notes[2] = "Ericsson"
	notes[3] = "BMAS"
	notes[4] = "Automation"

	fmt.Println("Notes :", notes)

	var dates [3]time.Time
	dates[0] = time.Now()
	dates[1] = time.Unix(1447920000, 0)
	dates[2] = time.Unix(1508632200, 0)

	fmt.Println("Dates : ", dates)

	var notes1 [5]string = [5]string{"do", "re", "mi", "fa", "so"}
	fmt.Println("notes1 : ", notes1)
	fmt.Println(notes1[0], notes1[2], notes1[4])

	for index, note := range notes1 {

		if index == 3 {
			notes1[3] = "Hello"
			fmt.Println("Updated array : ", notes1)
		}

		fmt.Println(index, " : ", note)
	}

	fmt.Println("Updated array : ", notes1, " with length : ", len(notes1))

}
