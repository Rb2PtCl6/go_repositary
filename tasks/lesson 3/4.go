package main

import (
	"fmt"
)

func main() {
	var a, b float64
	fmt.Print("a: ")
	fmt.Scan(&a)
	fmt.Println()
	fmt.Print("b: ")
	fmt.Scan(&b)
	fmt.Println()
	if a != 0 {
		fmt.Println(-b / a)
	} else if a == 0 && b != 0 {
		fmt.Println("Корней нет")
	} else if a == 0 && b == 0 {
		fmt.Println("Любое число")
	} else {
		fmt.Println("Unknown error!")
	}

}
