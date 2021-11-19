//Package date holds structure for date creation and methods for date creation and query!
package date

import (
	"errors"
	"fmt"
)

type Date struct {
	year  int
	month int
	day   int
}

func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid input for year")
	}
	d.year = year
	fmt.Println("SetYear called:", d.year)
	return nil
}

func (d *Date) SetMonth(month int) error {
	if month < 1 && month > 12 {
		return errors.New("invalid input for month")
	}
	d.month = month
	fmt.Println("SetMonth called:", d.month)
	return nil
}

func (d *Date) SetDay(day int) error {
	if day < 1 && day > 31 {
		return errors.New("invalid input for day")
	}
	d.day = day
	fmt.Println("SetDay called:", d.day)
	return nil
}

func (d *Date) Year() int {
	fmt.Println("Queried Year :", d.year)
	return d.year
}

func (d *Date) Month() int {
	fmt.Println("Queried Month :", d.month)
	return d.month
}

func (d *Date) Day() int {
	fmt.Println("Queried Day :", d.day)
	return d.day
}
