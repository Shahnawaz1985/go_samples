package main

import (
	"fmt"
	"sort"
)

func main() {
	var userdetail map[string]string
	userdetail = make(map[string]string)

	userdetail["Name"] = "Automation Architect"
	userdetail["JobPosition"] = "Job Stage 6"
	userdetail["BusinessUnit"] = "BMAS Q&OE"
	userdetail["CurrentAssignment"] = "MBNL Hardy Transformation"

	fmt.Println("userdetail :", userdetail)

	fmt.Printf("%#v\n:", userdetail)

	detail, isFound := userdetail["Salary"]

	if isFound {
		fmt.Println("Identified value for detail : ", detail)
	} else {
		fmt.Println("no value found!")
	}

	for key, value := range userdetail {
		fmt.Println(key, ":", value)
	}

	var mapkeys []string

	for mapkey := range userdetail {
		mapkeys = append(mapkeys, mapkey)
	}

	sort.Strings(mapkeys)

	fmt.Println("***********************Map print with sorted keys!***************************")
	for _, sortedkey := range mapkeys {
		fmt.Println("Value for ", sortedkey, "is :", userdetail[sortedkey])
	}
}
