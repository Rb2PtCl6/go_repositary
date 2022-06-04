package main

import (
	"fmt"
)

func main() {
	var a, b, a1, b1 int
	rezP, rezN, rezE := "yes", "no", "Unknown error!"
	var rezF string
	fmt.Print("a: ")
	fmt.Scan(&a)
	fmt.Println()
	fmt.Print("b: ")
	fmt.Scan(&b)
	fmt.Println()
	fmt.Print("a1: ")
	fmt.Scan(&a1)
	fmt.Println()
	fmt.Print("b1: ")
	fmt.Scan(&b1)
	fmt.Println()
	if a > a1 && b > b1 {
		rezF = rezP
	} else if a < (a1/2) || b < (b1/2) {
		rezF = rezN
	} else if a > (a1 / 2) {
		rezF = rezP
	} else if b > (b1 / 2) {
		rezF = rezP
	} else {
		rezF = rezE
	}
	fmt.Println(rezF)
}
