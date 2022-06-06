package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Print("num: ")
	fmt.Scan(&num)
	fmt.Println()
	if ((num%2) == 0 && num != 2) || ((num%3) == 0 && num != 3) || (num%4) == 0 || ((num%5) == 0 && num != 5) || ((num%7) == 0 && num != 7) {
		fmt.Println("Число не является простым. ")
	} else {
		fmt.Println("Число является простым.")
	}
}
