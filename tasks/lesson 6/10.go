package main

import (
	"fmt"
)

func line(n int) {
	for i := 0; i < n; i++ {
		fmt.Print("x")
	}
	fmt.Println()
}
func draw10(n1, n2, n3, n4, n5 int) {
	line(n1)
	line(n2)
	line(n3)
	line(n4)
	line(n5)
}
func main() {
	var n1, n2, n3, n4, n5 int
	fmt.Print("n1: ")
	fmt.Scan(&n1)
	fmt.Println()
	fmt.Print("n2: ")
	fmt.Scan(&n2)
	fmt.Println()
	fmt.Print("n3: ")
	fmt.Scan(&n3)
	fmt.Println()
	fmt.Print("n4: ")
	fmt.Scan(&n4)
	fmt.Println()
	fmt.Print("n5: ")
	fmt.Scan(&n5)
	fmt.Println()
	draw10(n1, n2, n3, n4, n5)
}
