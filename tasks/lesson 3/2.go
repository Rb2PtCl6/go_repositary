package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	rezP, rezN, rezE := "exist", "not exist", "Unknown error!"
	var rezF string
	fmt.Print("a: ")
	fmt.Scan(&a)
	fmt.Println()
	fmt.Print("b: ")
	fmt.Scan(&b)
	fmt.Println()
	fmt.Print("c: ")
	fmt.Scan(&c)
	fmt.Println()
	if (a + b) > c {
		rezF = rezP
	} else if (a + c) > b {
		rezF = rezP
	} else if (c + b) > a {
		rezF = rezP
	} else if a > (c + b) {
		rezF = rezN
	} else if b > (a + c) {
		rezF = rezN
	} else if c > (a + b) {
		rezF = rezN
	} else {
		rezF = rezE
	}
	fmt.Println(rezF)
}
