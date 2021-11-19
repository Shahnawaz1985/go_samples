package customtype

import (
	"date"
	"errors"
	"fmt"
	"unicode/utf8"
)

type ExtendedUser struct {
	name     string
	age      float64
	userid   string
	password string
	Address
	date.Date
}

type Address struct {
	addressLine1 string
	addressLine2 string
	landMark     string
	pincode      float64
	city         string
	state        string
	country      string
	isActive     bool
}

func (u *ExtendedUser) SetName(name string) error {
	if utf8.RuneCountInString(name) > 35 {
		return errors.New("invalid input for name size > 35")
	}
	u.name = name
	fmt.Println("SetName called:", u.name)
	return nil
}

func (u *ExtendedUser) SetAge(age float64) error {
	if recursionCountDigits(age) > 2 {
		return errors.New("invalid age input with 3 digit value")
	}
	u.age = age
	fmt.Println("SetAge called:", u.age)
	return nil
}

func (u *ExtendedUser) Setuserid(userID string) error {
	if utf8.RuneCountInString(userID) > 15 {
		return errors.New("invalid input for userid size > 15")
	}
	u.userid = userID
	fmt.Println("SetUserid called:", u.userid)
	return nil
}

func (u *ExtendedUser) SetPassword(password string) error {
	if utf8.RuneCountInString(password) > 20 {
		return errors.New("invalid input for password size > 20")
	}
	u.password = password
	fmt.Println("SetPassword called:", u.password)
	return nil
}

func (a *Address) SetAddressLine1(addrLine1 string) error {
	if addrLine1 == "" {
		return errors.New("invalid input for addrLine1-blank input")
	}
	a.addressLine1 = addrLine1
	fmt.Println("SetAddressLine1 called:", a.addressLine1)
	return nil
}

func (a *Address) SetAddressLine2(addrLine2 string) error {
	if addrLine2 == "" {
		return errors.New("invalid input for addrLine2-blank input")
	}
	a.addressLine2 = addrLine2
	fmt.Println("SetAddressLine2 called:", a.addressLine2)
	return nil
}

func (a *Address) SetLandMark(lndMrk string) error {
	if lndMrk == "" {
		return errors.New("invalid input for lndMrk-blank input")
	}
	a.landMark = lndMrk
	fmt.Println("SetLandMark called:", a.landMark)
	return nil
}

func (a *Address) SetPincode(pncd float64) error {
	if recursionCountDigits(pncd) == 0 {
		return errors.New("invalid input for pncd-blank input")
	}
	a.pincode = pncd
	fmt.Println("SetPincode called:", a.pincode)
	return nil
}

func (a *Address) SetCity(city string) error {
	if city == "" {
		return errors.New("invalid input for city-blank input")
	}
	a.city = city
	fmt.Println("SetCity called:", a.city)
	return nil
}

func (a *Address) SetState(stt string) error {
	if stt == "" {
		return errors.New("invalid input for stt-blank input")
	}
	a.state = stt
	fmt.Println("SetState called:", a.state)
	return nil
}

func (a *Address) SetCountry(cntry string) error {
	if cntry == "" {
		return errors.New("invalid input for cntry-blank input")
	}
	a.country = cntry
	fmt.Println("SetCountry called:", a.country)
	return nil
}

func (a *Address) SetIsActive(active bool) error {
	a.isActive = active
	fmt.Println("SetIsActive called:", a.isActive)
	return nil
}

func (u *ExtendedUser) CustomPrintStruct(s *ExtendedUser) {
	fmt.Println("Name : ", s.name, ", Age :", s.age, ", Userid :", s.userid, ", Password :", s.password)
	fmt.Println("CommunicationAddress - AddressLine1 :", s.addressLine1, ", AddressLine2 :", s.addressLine2, ", "+
		"LandMark :", s.landMark, ", Pincode :", s.pincode, ", City :", s.city, ", State :", s.state, ", "+
		"Country :", s.country, ", IsActive :", s.isActive)
}

func recursionCountDigits(number float64) int {
	if number < 10 {
		return 1
	} else {
		return 1 + recursionCountDigits(number/10)
	}
}
