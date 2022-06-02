package main

import (
	"fmt"
)

func main() {
	var num float64
	fmt.Print("num: ")
	fmt.Scan(&num)
	fmt.Println()
	if num < 0 {
		num *= (-1)
	}
	fmt.Println(num)
}
