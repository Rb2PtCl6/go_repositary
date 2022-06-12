package main

import (
	"fmt"
	"math"
)

func elem(n []int, numTo int) (n1 []int) {
	n1 = make([]int, numTo)
	i := 0
	for ; i < len(n); i++ {
		n1[i] = n[i]
	}
	for ; i < numTo; i++ {
		n1[i] = int(math.Max(float64(n1[i-2]), float64(n1[i-5])))
	}
	return
}
func main() {
	var numTo int
	n := make([]int, 5)
	fmt.Print("numTo: ")
	fmt.Scan(&numTo)
	fmt.Println()
	for i := 0; i < 5; i++ {
		fmt.Print("slice: ")
		fmt.Scan(&n[i])
		fmt.Println()
	}
	rez := elem(n, numTo)
	fmt.Println(rez)
}
