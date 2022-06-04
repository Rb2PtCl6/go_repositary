package main

import (
	"fmt"
)

func main() {
	var a, b, c float64
	//a
	fmt.Print("a: ")
	fmt.Scan(&a)
	fmt.Println()
	//b
	fmt.Print("b: ")
	fmt.Scan(&b)
	fmt.Println()
	//c
	fmt.Print("c: ")
	fmt.Scan(&c)
	fmt.Println()
	// logic
	if a > b && a > c && b > c {
		fmt.Println(a)
		fmt.Println(b)
		fmt.Println(c)
	} else if a > b && a > c && b < c {
		fmt.Println(a)
		fmt.Println(c)
		fmt.Println(b)
	} else if b > a && b > c && a > c {
		fmt.Println(b)
		fmt.Println(a)
		fmt.Println(c)
	} else if b > a && b > c && a < c {
		fmt.Println(b)
		fmt.Println(c)
		fmt.Println(a)
	} else if c > a && c > b && a > b {
		fmt.Println(c)
		fmt.Println(a)
		fmt.Println(b)
	} else if c > a && c > b && a < b {
		fmt.Println(c)
		fmt.Println(b)
		fmt.Println(a)
	} else {
		fmt.Println("Unknown error!")
	}
}
