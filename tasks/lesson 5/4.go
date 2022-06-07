package main

import (
	"fmt"
	"math"
)

func lessmore(n float64) (float64, float64) {
	nS := math.Sqrt(n)
	perc := (nS / 10) / 2
	less := math.Round(nS - perc)
	more := math.Round(nS + perc)
	return less, more
}
func main() {
	var n float64
	fmt.Print("n: ")
	fmt.Scan(&n)
	l, m := lessmore(n)
	fmt.Println(l)
	fmt.Println(m)
}
