package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num, lengthS int
	var rez string
	var rez2 int = 0
	fmt.Print("num: ")
	fmt.Scan(&num)
	fmt.Println()
	fmt.Print("length: ")
	fmt.Scan(&lengthS)
	fmt.Println()
	for i := lengthS; i > 0; i-- {
		rez += strconv.Itoa((num % 10))
		num /= 10
	}
	for i1 := lengthS; i1 > 0; i1-- {
		rez2 += (num % 10)
		num /= 10
		rez2 *= 10
	}
	fmt.Println(rez)
	fmt.Println(rez2)
}
