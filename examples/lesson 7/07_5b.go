package main

import "fmt"

func digits(a uint) (n [10]int) {
	n[a%10]++
	for a /= 10; a != 0; a /= 10 {
		n[a%10]++
	}
	return
}

func main() {
	var a uint
	fmt.Scan(&a)
	fmt.Println(digits(a))
}
