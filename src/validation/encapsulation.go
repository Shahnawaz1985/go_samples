package main

import (
	"date"
	"fmt"
	"log"
)

func main() {
	date_tmp := date.Date{}

	err := date_tmp.SetYear(2021)
	if err != nil {
		log.Fatal(err)
	}

	err = date_tmp.SetMonth(11)
	if err != nil {
		log.Fatal(err)
	}

	err = date_tmp.SetDay(16)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(date_tmp.Year())
	fmt.Println(date_tmp.Month())
	fmt.Println(date_tmp.Day())
}
