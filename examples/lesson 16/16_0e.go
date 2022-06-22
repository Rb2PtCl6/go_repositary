package main

import (
	"fmt"
)

type (
	tPerson struct {
		name      tFullName
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

	pers := tPerson{
		name:      tFullName{firstName: "Greg", lastName: "Lee"},
		salary:    3500,
		birthDate: tDate{1991, 8, 19},
	}
	fmt.Println(pers)

	fmt.Print(pers.name.firstName, " ", pers.name.lastName)
	fmt.Print(" was born on ")
	fmt.Println(pers.birthDate.day, "/", pers.birthDate.month, "/", pers.birthDate.year)

	bd := tDate{
		day:   8,
		month: 7,
		year:  1983,
	}
	var pers2 tPerson
	pers2.birthDate = bd
	pers2.salary = 4250
	pers2.name = tFullName{"Basil", "Pupkin"}
	fmt.Println(pers2)
}
