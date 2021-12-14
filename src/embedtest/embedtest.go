package main

import (
	_ "embed"
	"fmt"
)

//go:embed test.txt
var fileContent string

func main() {
	fmt.Println(fileContent)
}
