package main

import (
	"fmt"
)

func draw1(n int) {
	line_n := 1
	num_s := 1
	for i := 0; i < n; i++ {
		for i1 := 0; i1 < line_n; i1++ {
			fmt.Printf("%v ", num_s)
			num_s++
			if i1 == line_n-1 {
				fmt.Println()
			}
		}
		line_n++
	}
}
func main() {
	var n int
	fmt.Print("n: ")
	fmt.Scan(&n)
	fmt.Println()
	draw1(n)
}
