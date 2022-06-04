package main

import (
	"fmt"
)

func main() {
	var x, y int
	kv1, kv2, kv3, kv4 := "1", "2", "3", "4"
	minX, plusX, minY, plusY := "-x", "+x", "-y", "+y"
	zero, rezE := "0;0", "Unknown error!"
	var rez string
	fmt.Print("x: ")
	fmt.Scan(&x)
	fmt.Println()
	fmt.Print("y: ")
	fmt.Scan(&y)
	fmt.Println()
	if x == 0 && y == 0 {
		rez = zero
	} else if x > 0 && y > 0 {
		rez = kv1
	} else if x > 0 && y < 0 {
		rez = kv2
	} else if x < 0 && y < 0 {
		rez = kv3
	} else if x < 0 && y > 0 {
		rez = kv4
	} else if x < 0 && y == 0 {
		rez = minX
	} else if x > 0 && y == 0 {
		rez = plusX
	} else if x == 0 && y < 0 {
		rez = minY
	} else if x == 0 && y > 0 {
		rez = plusY
	} else {
		rez = rezE
	}
	fmt.Println(rez)
}
