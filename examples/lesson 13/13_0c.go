package main

import (
	"fmt"
)

func main() {
	var personSalary map[string]int

	if personSalary == nil {
		//  The `if' condition is always true!
		fmt.Println("map is nil. Going to make one.")
		personSalary = make(map[string]int)
	}

	personSalary["steve"] = 12_000
	personSalary["jamie"] = 15_000
	personSalary["mike"] = 9_000
	fmt.Println("personSalary map contents:", personSalary)

	var otherPersonList map[string]int
	otherPersonList = personSalary

	personSalary["mike"] = 10_000
	fmt.Println("otherPersonList map contents:", otherPersonList)

	personSalary["john"] = 10_500
	fmt.Println("otherPersonList map contents:", otherPersonList)

	delete(personSalary, "jamie")
	fmt.Println("otherPersonList map contents:", otherPersonList)
}

/*
map is nil. Going to make one.
personSalary map contents: map[jamie:15000 mike:9000 steve:12000]
otherPersonList map contents: map[jamie:15000 mike:10000 steve:12000]
otherPersonList map contents: map[jamie:15000 john:10500 mike:10000 steve:12000]
otherPersonList map contents: map[john:10500 mike:10000 steve:12000]
*/ 
