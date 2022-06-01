package main

import (
	"fmt"
)

func main() {
	var num, pak2, pak3, pak4 float64
	fmt.Print("num: ")
	fmt.Scan(&num)
	pak2 = num * num
	pak3 = pak2 * num
	pak4 = pak3 * num
	fmt.Print("a in 1: ", num)
	fmt.Println()
	fmt.Print("a in 2: ", pak2)
	fmt.Println()
	fmt.Print("a in 3: ", pak3)
	fmt.Println()
	fmt.Print("a in 4: ", pak4)
}
