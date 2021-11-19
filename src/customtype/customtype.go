package customtype

import (
	"fmt"
)

type CustomType string

type CustomTypeInt int

type Number float64

func (c CustomType) SayCustomHi() {
	fmt.Println("SayCustomHi from ", c)
}

func (c CustomType) CustomCallWithParams(number int, flag bool) CustomTypeInt {
	fmt.Println("CustomType :", c)
	fmt.Println("number :", number)
	fmt.Println("flag :", flag)
	return 10
}

func (n *Number) DoubleBy2() Number {
	return *n * 2
}
