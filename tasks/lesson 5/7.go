package main

import (
	"fmt"
)

func y_vis(day float64) int {
	year := day / (365*4 + 1)
	year = 4 * year
	year = year + 2001
	return int(year)
}
func y_nevis(day float64) int {
	year := day / 365
	year = year + 2001
	return int(year)
}
func main() {
	var day float64
	fmt.Print("day: ")
	fmt.Scan(&day)
	fmt.Println()
	vis := y_vis(day)
	nevis := y_nevis(day)
	fmt.Printf("vis: %v\n", vis)
	fmt.Printf("nevis: %v\n", nevis)
}
