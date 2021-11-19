package main

import (
	"fileiterator_test"
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

	defer fileiterator_test.ReportPanic()
	fileiterator_test.ScanDirectory("/Users/mohammadshahnawazakhter/intellij_go/GO_SAMPLE_PROJECT")

	err := fmt.Errorf("throwing error intentionally")
	panic(err)
}
