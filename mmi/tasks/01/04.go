package main

import (
	"fmt"
	"math"
)

func main() {
	var rad, sO, around float64
	fmt.Print("rad: ")
	fmt.Scan(&rad)
	fmt.Println()
	sO = math.Pi * rad * rad
	around = 2 * math.Pi * rad
	fmt.Print("sO: ", sO)
	fmt.Println()
	fmt.Print("around: ", around)
	fmt.Println()
}
