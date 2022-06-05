package main

import (
	"fmt"
)

func main() {
	var n, m int // n-blue, m-red
	var sum int
	fmt.Print("n: ")
	fmt.Scan(&n)
	fmt.Println()
	fmt.Print("m: ")
	fmt.Scan(&m)
	fmt.Println()
	sum = n + m
	n1 := float64(n)
	m1 := float64(m)
	sum1 := float64(sum)
	fmt.Println(n1 / sum1)
	fmt.Println(m1 / sum1)
	fmt.Println((n1 / sum1) * (m1 / (sum1 - 1)))
	fmt.Println((m / sum) * (m - 1/sum - 1) * (m - 2/sum - 2) * (m - 3/sum - 3) * (m - 4/sum - 4))
	// как реализовать этот пункт без for?
}
