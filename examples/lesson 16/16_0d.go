package main

import (
	"fmt"
)

type tEmployee struct {
	firstName, lastName string
	age                 int
	salary              int
}

func main() {
	//creating structure using field names
	emp := tEmployee{
		firstName: "Sam",
		age:       25,
		salary:    2500,
		lastName:  "Anderson",
	}
	fmt.Println(emp)

	//assignment structure without using field names
	emp = tEmployee{"Thomas", "Paul", 29, 2800}
	fmt.Println(emp)

	emp2 := struct {
		name string
		age  int
	}{
		age:  33,
		name: "John",
	}
	fmt.Println(emp2)

	emp2 = struct {
		name string
		age  int
	}{"Bill", 63}
	fmt.Println(emp2)
}
