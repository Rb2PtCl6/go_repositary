package main

import (
	"fmt"
)

func first_line(n int) {
	fmt.Print("xxx")
	for i := 2; i <= n; i++ {
		fmt.Print("xx")
	}
	fmt.Println()
}
func second_line(n int) {
	fmt.Print("x x")
	for i := 2; i <= n; i++ {
		fmt.Print(" x")
	}
	fmt.Println()
}
func draw3(n, m int) {
	for i := 0; i < m; i++ {
		first_line(n)
		second_line(n)
	}
	first_line(n)
	fmt.Println()
}
func main() {
	var n, m int
	fmt.Print("n: ")
	fmt.Scan(&n)
	fmt.Println()
	fmt.Print("m: ")
	fmt.Scan(&m)
	fmt.Println()
	draw3(n, m)
}
