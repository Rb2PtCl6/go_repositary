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
	var emp tEmployee
	emp.firstName = "Sam"
	emp.lastName = "Anderson"
	emp.age, emp.salary = 55, 6000
	fmt.Println(emp)
}
