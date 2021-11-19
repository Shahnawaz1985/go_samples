package generictypes

import (
	"fmt"
)

type User struct {
	Name                 string
	Age                  float64
	Userid               string
	Password             string
	CommunicationAddress Address
}

type Address struct {
	AddressLine1 string
	AddressLine2 string
	LandMark     string
	Pincode      float64
	City         string
	State        string
	Country      string
	IsActive     bool
}

func PrintStruct(s *User) {
	fmt.Println("Name : ", s.Name, ", Age :", s.Age, ", Userid :", s.Userid, ", Password :", s.Password)
	fmt.Println("CommunicationAddress - AddressLine1 :", s.CommunicationAddress.AddressLine1, ", AddressLine2 :", s.CommunicationAddress.AddressLine2, ", "+
		"LandMark :", s.CommunicationAddress.LandMark, ", Pincode :", s.CommunicationAddress.Pincode, ", City :", s.CommunicationAddress.City, ", State :", s.CommunicationAddress.State, ", "+
		"Country :", s.CommunicationAddress.Country, ", IsActive :", s.CommunicationAddress.IsActive)
}

func PopulateDefaultUser(s *User) *User {
	fmt.Println("************Creating Default User************")
	s.Userid = "test-User-2"
	s.Name = "Generic User"
	s.Age = 25
	s.Password = "test-pwd"
	s.CommunicationAddress.AddressLine1 = "First Lane - ABC"
	s.CommunicationAddress.AddressLine2 = "Second Lane - DEF"
	s.CommunicationAddress.City = "New York"
	s.CommunicationAddress.State = "North Carolina"
	s.CommunicationAddress.Country = "United States"
	s.CommunicationAddress.IsActive = true
	s.CommunicationAddress.LandMark = "Downtown Hill 6"
	s.CommunicationAddress.Pincode = 64006
	PrintStruct(s)
	fmt.Println("************Default User Created************")
	return s
}
