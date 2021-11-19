package main

import (
	"fmt"
)

func catchError() {
	t := recover()
	fmt.Println("recover() called.")
	err, ok := t.(error)
	if ok {
		fmt.Println(err.Error())
	}

}

func main() {
	defer catchError()
	err := fmt.Errorf("throwing error intentionally")
	panic(err)
}
