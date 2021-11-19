package main

import (
	"customtype"
	"fmt"
)

func main() {

	customTp := customtype.ExtendedUser{}
	fmt.Println("Setting up custom type for encapsulation test!")
	fmt.Println("**************************************************")

	err := customTp.SetName("Shahnawaz Akhter")
	if err != nil {
		fmt.Println("Attribute set successful.")
	}
	err = customTp.SetAge(35)
	if err != nil {
		fmt.Println("Attribute set successful.")
	}
	err = customTp.Setuserid("XXX111AAA")
	if err != nil {
		fmt.Println("Attribute set successful.")
	}
	err = customTp.SetPassword("Test123456")
	if err != nil {
		fmt.Println("Attribute set successful.")
	}
	err = customTp.SetAddressLine1("TestAddressLine1 set")
	if err != nil {
		fmt.Println("Attribute set successful.")
	}
	err = customTp.SetAddressLine2("TestAddressLine2 set")
	if err != nil {
		fmt.Println("Attribute set successful.")
	}
	err = customTp.SetLandMark("landmark set")
	if err != nil {
		fmt.Println("Attribute set successful.")
	}
	err = customTp.SetPincode(112233)
	if err != nil {
		fmt.Println("Attribute set successful.")
	}
	err = customTp.SetCity("TestCity set")
	if err != nil {
		fmt.Println("Attribute set successful.")
	}
	if err != nil {
		fmt.Println("Attribute set successful.")
	}
	err = customTp.SetState("TestState set")
	if err != nil {
		fmt.Println("Attribute set successful.")
	}
	err = customTp.SetCountry("TestCountry set")
	if err != nil {
		fmt.Println("Attribute set successful.")
	}
	err = customTp.SetIsActive(true)
	if err != nil {
		fmt.Println("Attribute set successful.")
	}
	customTp.CustomPrintStruct(&customTp)

	fmt.Println("**************************************************")

}
