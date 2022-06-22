package main

import (
	"fmt"
)

type (
	tPerson struct {
		tFullName
		birthDate tDate
		salary    int
	}
	tDate struct {
		year  int
		month int
		day   int
	}
	tFullName struct {
		firstName, lastName string
	}
)

func main() {
	bd := tDate{
		day:   8,
		month: 7,
		year:  1983,
	}
	var pers2 tPerson
	pers2.birthDate = bd
	pers2.salary = 4250
	pers2.firstName = "Basil"
	pers2.lastName = "Pupkin"
	fmt.Println(pers2)

	// It's impossible to declare a first and last name in this
	// statement because the corresponfing field is promoted (anonymous)
	pers := tPerson{
		salary:    3500,
		birthDate: tDate{1991, 8, 19},
	}
	// although we can refer to the fields firstName and lastName name
	pers.firstName = "Greg"
	pers.lastName = "Lee"
	fmt.Println(pers)

	fmt.Print(pers.firstName, " ", pers.lastName)
	fmt.Print(" was born on ")
	fmt.Println(pers.birthDate.day, "/", pers.birthDate.month, "/", pers.birthDate.year)

}
