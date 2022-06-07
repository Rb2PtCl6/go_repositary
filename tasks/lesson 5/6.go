package main

import (
	"fmt"
)

func delit2(num int) {
	var i int = 1
	for {
		if num%i == 0 && num != i {
			fmt.Println(i)
		} else if num == i {
			fmt.Println(i)
			break
		}
		i++
	}
}
func main() {
	var num int
	fmt.Print("num: ")
	fmt.Scan(&num)
	fmt.Println()
	delit2(num)
}
