package main

import (
	"fmt"
)

func main() {
	var day int
	wn, sp, sm, au, rezE := "winter", "spring", "summer", "autumn", "Unknown error!"
	var rez string
	fmt.Print("day: ")
	fmt.Scan(&day)
	fmt.Println()
	// январь-wn-31
	// февраль-wn-28
	// март-sp-31
	// апрель-sp-30
	// май-sp-31
	// июнь-sm-30
	// июль-sm-31
	// август-sm-31
	// сентябрь-au-30
	// октябрь-au-31
	// ноябрь-au-30
	// декабрь-wn-31
	if day <= 59 || (day <= 365 && day >= 335) {
		rez = wn
	} else if day >= 60 && day <= 152 {
		rez = sp
	} else if day >= 153 && day <= 244 {
		rez = sm
	} else if day >= 245 && day <= 334 {
		rez = au
	} else {
		rez = rezE
	}
	fmt.Println(rez)
}
