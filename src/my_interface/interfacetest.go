package my_interface

import (
	"fmt"
)

type MyInterface interface {
	MethodWithoutParameters()
	MethodWithParameter(float64)
	MethodWithReturnValue() string
}

type MyTestType string

func (m MyTestType) MethodWithoutParameters() {
	fmt.Println("MethodWithoutParameters called for :", m)
}

func (m MyTestType) MethodWithParameter(t float64) {
	fmt.Println("MethodWithParameters called for :", t, " on type :", m)
}

//MethodWithReturnValue is used for returning values.%#v
func (m MyTestType) MethodWithReturnValue() string {
	return "MethodWithReturnValue called."
}

func (m MyTestType) MethodNotInInterface() {
	fmt.Println("MethodNotInInterface called for :", m)
}
