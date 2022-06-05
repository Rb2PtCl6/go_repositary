package main

import (
	"fmt"
)

func main() {
	var ord, size int
	rezE := "Unknown error!"
	fmt.Println("Выберите доску:")
	fmt.Println("	1) 8x8 ")
	fmt.Println("	2) 1000x1000 ")
	fmt.Scan(&ord)
	fmt.Println()
	if ord == 1 {
		size = 8
	} else if ord == 2 {
		size = 1000
	} else {
		fmt.Println(rezE)
	}
	if size == 8 || size == 1000 {
		var a, b int
		var rez string
		bl, wt := "black", "white"
		fmt.Print("a: ")
		fmt.Scan(&a)
		fmt.Println()
		fmt.Print("b: ")
		fmt.Scan(&b)
		fmt.Println()
		if a <= size && b <= size {
			// (--) (a1) (a2) (a3) (a4)
			// (b1)  bl   wt   bl   wt
			// (b2)  wt   bl   wt   bl
			// (b3)  bl   wt   bl   wt
			// (b4)  wt   bl   wt   bl
			if (a%2) == 1 && (b%2) == 1 {
				rez = bl //a1-b1
			} else if (a%2) == 0 && (b%2) == 0 {
				rez = bl //a2-b2
			} else if (a%2) == 0 && (b%2) == 1 {
				rez = wt
			} else if (a%2) == 1 && (b%2) == 0 {
				rez = wt
			}
		} else {
			rez = rezE
		}
		fmt.Println(rez)
	}
}
