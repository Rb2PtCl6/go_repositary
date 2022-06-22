package main

import (
	"fmt"
)

func main() {
	var month int
	fmt.Print("Введите номер месяца: ")
	fmt.Scan(&month)
	switch month {
	case 3, 4, 5:
		fmt.Println("весна")
	case 6, 7, 8:
		fmt.Println("лето")
	case 9, 10, 11:
		fmt.Println("осень")
	case 12, 1, 2:
		fmt.Println("зима")
	}
}
