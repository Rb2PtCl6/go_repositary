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
	emp.lastName = "Anderson"
	emp.salary = 6000
	fmt.Println(emp)
}
