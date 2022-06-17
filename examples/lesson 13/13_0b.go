package main

import (
	"fmt"
)

func main() {
	personSalary := make(map[string]int)
	personSalary["steve"] = 12_000
	personSalary["jamie"] = 15_000
	personSalary["mike"] = 9_000
	fmt.Printf("personSalary map contents: %v\n", personSalary)
	employee := "jamie"
	fmt.Printf("Salary of %s is %d\n", employee, personSalary[employee])
	fmt.Println("Salary of joe is", personSalary["joe"])

	newEmp := "joe"
	if value, ok := personSalary[newEmp]; ok {
		fmt.Printf("Salary of %v is %v\n", newEmp, value)
	} else {
		fmt.Println(newEmp, "not found")
	}

	fmt.Println(len(personSalary))		
	delete(personSalary, "steve")
	fmt.Println(len(personSalary), personSalary)

	delete(personSalary, "john")
	fmt.Println(len(personSalary), personSalary)

	total := 0
	for employee, salary := range personSalary {
		fmt.Printf("%s - %d\n", employee, salary)
		total += salary
	}
	fmt.Println(total)
}

/*
personSalary map contents: map[jamie:15000 mike:9000 steve:12000]
Salary of jamie is 15000
Salary of joe is 0
joe not found
3
2 map[jamie:15000 mike:9000]
2 map[jamie:15000 mike:9000]
jamie - 15000
mike - 9000
24000
*/ 
