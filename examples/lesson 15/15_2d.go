package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Enter integer: ")
	fmt.Scan(&n)
	switch {
	case n < 0:
		fmt.Println("Negative")
	case n == 0:
		fmt.Println("Zero")
	case n <= 100:
		fmt.Println("Small")
	default:
		fmt.Println("Large")
	}
}
