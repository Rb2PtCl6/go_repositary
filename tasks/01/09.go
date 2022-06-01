package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2, yS, xS, att float64
	fmt.Print("x1: ")
	fmt.Scan(&x1)
	fmt.Println()
	fmt.Print("y1: ")
	fmt.Scan(&y1)
	fmt.Println()
	fmt.Print("x2: ")
	fmt.Scan(&x2)
	fmt.Println()
	fmt.Print("y2: ")
	fmt.Scan(&y2)
	fmt.Println()
	yS = x2 - x1
	xS = x2 - x1
	att = math.Sqrt((yS * yS) + (xS * xS))
	fmt.Print("att: ", att)
}
