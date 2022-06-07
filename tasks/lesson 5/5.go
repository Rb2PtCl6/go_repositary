package main

import (
	"fmt"
	"math"
)

func area(a, b, c float64) float64 {
	if a+b < c || a+c < b || c+b < a {
		return -1
	}
	p := (a + b + c) / 2
	area := math.Sqrt(p * (p - a) * (p - b) * (p - c))
	return area
}
func main() {
	var a, b, c float64
	fmt.Print("a: ")
	fmt.Scan(&a)
	fmt.Println()
	fmt.Print("b: ")
	fmt.Scan(&b)
	fmt.Println()
	fmt.Print("c: ")
	fmt.Scan(&c)
	fmt.Println()
	fmt.Println(area(a, b, c))
}

//
