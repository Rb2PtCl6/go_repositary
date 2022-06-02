package main

import (
	"fmt"
)

func main() {
	var a, b float64
	var rez1, rez2, rez3, rez4 float64
	//a
	fmt.Print("a: ")
	fmt.Scan(&a)
	fmt.Println()
	//b
	fmt.Print("b: ")
	fmt.Scan(&b)
	fmt.Println()
	//logic
	rez1 = a + b
	rez2 = a - b
	rez3 = a * b
	if b != 0 {
		rez4 = a / b
	}
	//compearing
	//
}
