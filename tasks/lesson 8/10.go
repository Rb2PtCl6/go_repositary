package main

import (
	"fmt"
)

func double(num []int) {
	num1 := make([]int, len(num)*2)
	i := 0
	for i < len(num) {
		num1[i*2], num1[i*2+1] = num[i], num[i]
		fmt.Println(num1[i*2])
		fmt.Println(num1[i*2+1])
		i++
	}
}
func main() {
	num := make([]int, 10)
	for i := 0; i < len(num); i++ {
		fmt.Printf("Number %v: ", i+1)
		fmt.Scan(&num[i])
		fmt.Println()
	}
	double(num)
}
