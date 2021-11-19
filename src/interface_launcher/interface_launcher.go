package main

import (
	"fmt"
	"my_interface"
)

func main() {
	var testValue my_interface.MyInterface
	testValue = my_interface.MyTestType("Test Type created.")
	fmt.Println("MyType created for :", testValue)

	testValue.MethodWithoutParameters()
	testValue.MethodWithParameter(150.6)
	fmt.Println(testValue.MethodWithReturnValue())

	testType := my_interface.MyTestType("Custom type for non-interface method")
	testType.MethodNotInInterface()

	fmt.Println("MethodNotInInterface calling with non-interface type.")

	myvar, ok := testValue.(my_interface.MyTestType)

	if ok {
		fmt.Println("MyTestType is captured as type assertion")
		myvar.MethodNotInInterface()
	}
	fmt.Printf("MethodNotInInterface called again for : %#v\n", myvar)
	fmt.Println("Test steps included.")
}
