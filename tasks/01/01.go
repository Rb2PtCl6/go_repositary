package main

import (
	"fmt"
)

func main() {
	var h, w, s float64
	fmt.Print("h: ")
	fmt.Scan(&h)
	fmt.Println()
	fmt.Print("w: ")
	fmt.Scan(&w)
	fmt.Println()
	s = h * w
	fmt.Print("Area: ", s)
}
