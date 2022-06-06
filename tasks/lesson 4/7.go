package main

import (
	"fmt"
)

func main() {
	var a, b float64
	a, b = 0, 1
	fmt.Println("0")
	for {
		a, b = b, a+b
		fmt.Println(a)
		if a >= 10*10*10*10*10*10 || b >= 10*10*10*10*10*10 {
			break
		}
	}
}
