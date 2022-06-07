package main

import (
	"fmt"
)

func check1(num int) bool {
	var i int = 1
	var rez int
	for {
		rez = i * i * i
		if rez == num {
			return true
		} else if rez > num {
			return false
		}
		i++
	}
}
func main() {
	var num int
	fmt.Print("num: ")
	fmt.Scan(&num)
	fmt.Println()
	fmt.Printf("Cube: %v\n", check1(num))
}
