package main

import (
	"fmt"
	"os"
	"variadictest"
)

func main() {
	teststring := []string{"Hello$", "World$", "TestContent$", "Array$", "Variadic$"}

	variadictest.TestVariadic(teststring...)

	teststring2 := []string{"1", "2", "3"}
	variadictest.TestVariadic(teststring2...)

	fmt.Println("Command line arguments:")
	fmt.Println(os.Args[1:])
}
