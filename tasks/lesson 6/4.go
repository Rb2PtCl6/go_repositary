package main

import (
	"fmt"
)

func draw4(n, m int) {
	for i := 0; i < n; i++ {
		for i1 := 0; i1 < m; i1++ {
			fmt.Print("x")
			if i1 == m-1 {
				fmt.Println()
			}
		}
	}
}
func main() {
	var n, m int
	fmt.Print("n: ")
	fmt.Scan(&n)
	fmt.Println()
	fmt.Print("m: ")
	fmt.Scan(&m)
	fmt.Println()
	draw4(n, m)
}
