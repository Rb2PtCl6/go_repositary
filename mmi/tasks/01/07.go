package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, x1, x2 float64
	fmt.Print("a: ")
	fmt.Scan(&a)
	fmt.Println()
	fmt.Print("b: ")
	fmt.Scan(&b)
	fmt.Println()
	fmt.Print("c: ")
	fmt.Scan(&c)
	fmt.Println()
	x1 = ((((-1) * b) + (1 * math.Sqrt((b*b)-(4*a*c)))) / (2 * a))
	x2 = ((((-1) * b) + ((-1) * math.Sqrt((b*b)-(4*a*c)))) / (2 * a))
	fmt.Print("x1: ", x1)
	fmt.Println()
	fmt.Print("x2: ", x2)
	fmt.Println()
}
