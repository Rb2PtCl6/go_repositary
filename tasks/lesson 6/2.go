package main

import (
	"fmt"
)

func draw2(n int) {
	line_n := 1
	num_s := 1
	for i := 0; i < n; i++ {
		for i1 := 0; i1 < line_n; i1++ {
			num_s++
		}
		line_n++
	}
	num_s--
	line_n--
	for i := 0; i < n; i++ {
		for i1 := 0; i1 < line_n; i1++ {
			fmt.Printf("%v ", num_s)
			num_s--
		}
		fmt.Println()
		line_n--
	}
}
func main() {
	var n int
	fmt.Print("n: ")
	fmt.Scan(&n)
	fmt.Println()
	draw2(n)
}
