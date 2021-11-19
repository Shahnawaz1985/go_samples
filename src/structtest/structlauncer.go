package main

import (
	"customtype"
	"fmt"
	"generictypes"
	"reflect"
)

var mystruct struct {
	number   float64
	name     string
	isActive bool
}

type user struct {
	name        string
	userid      string
	password    string
	age         float64
	description string
}

type Liters float64
type Gallons float64

func main() {
	mystruct.number = 64
	mystruct.name = "Shahnawaz"
	mystruct.isActive = true

	fmt.Println("mystruct :", mystruct)

	var genericUser user

	genericUser.userid = "test-user"
	genericUser.password = "password"
	genericUser.name = "anyUser"
	genericUser.age = 20
	genericUser.description = "test-user created for testing type struct functionality!"

	fmt.Println("genericUser : ", genericUser)

	var testUser generictypes.User

	testUser.Userid = "test-User-2"
	testUser.Name = "Generic User"
	testUser.Age = 25
	testUser.Password = "test-pwd"
	testUser.CommunicationAddress.AddressLine1 = "First Lane - ABC"
	testUser.CommunicationAddress.AddressLine2 = "Second Lane - DEF"
	testUser.CommunicationAddress.City = "New York"
	testUser.CommunicationAddress.State = "North Carolina"
	testUser.CommunicationAddress.Country = "United States"
	testUser.CommunicationAddress.IsActive = true
	testUser.CommunicationAddress.LandMark = "Downtown Hill 6"
	testUser.CommunicationAddress.Pincode = 64006

	generictypes.PrintStruct(&testUser)
	fmt.Println("Type of testUser:", reflect.TypeOf(testUser))

	testUser2 := generictypes.PopulateDefaultUser(&testUser)

	generictypes.PrintStruct(testUser2)

	fmt.Println("Type of testUser2:", reflect.TypeOf(testUser2))

	var carFuel Gallons
	var busFuel Liters

	carFuel = Gallons(10.0)
	busFuel = Liters(25.0)

	fmt.Println("carFuel:", carFuel, ", busFuel :", busFuel)

	value := customtype.CustomType("CustomType Test")
	value.SayCustomHi()

	value2 := customtype.CustomType("CustomType with Param")
	fmt.Println("Calling CustomCallWithParams :", value2.CustomCallWithParams(24, true))

	number := customtype.Number(5)
	fmt.Println("Calling Double() wiht asterisk :", number.DoubleBy2())

}
