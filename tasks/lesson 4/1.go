package main

import (
	"fmt"
)

func main() {
	for {
		var cel, far, kel, r float64
		fmt.Print("cel: ")
		fmt.Scan(&cel)
		fmt.Println()
		far = (cel * 1.8) + 32
		kel = 273 + cel
		r = cel * 0.8
		fmt.Print("far: ", far)
		fmt.Println()
		fmt.Print("kel: ", kel)
		fmt.Println()
		fmt.Print("r: ", r)
		fmt.Println()
		fmt.Println("continue: y / n")
		var n string
		fmt.Scan(&n)
		fmt.Println()
		if n == "n" {
			break
		}
	}
}
