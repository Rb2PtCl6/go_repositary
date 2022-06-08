package main

import (
	"fmt"
)

func line2(a1, a2, a3, a4, a5 bool) {
	if a1 {
		fmt.Print(" x")
	} else {
		fmt.Print("  ")
	}
	if a2 {
		fmt.Print(" x")
	} else {
		fmt.Print("  ")
	}
	if a3 {
		fmt.Print(" x")
	} else {
		fmt.Print("  ")
	}
	if a4 {
		fmt.Print(" x")
	} else {
		fmt.Print("  ")
	}
	if a5 {
		fmt.Print(" x")
	} else {
		fmt.Print("  ")
	}
	fmt.Println()
}
func draw11(n1, n2, n3, n4, n5 int) {
	var max int
	max = n1
	if n2 > max {
		max = n2
	}
	if n3 > max {
		max = n3
	}
	if n4 > max {
		max = n4
	}
	if n5 > max {
		max = n5
	}
	var zn1, zn2, zn3, zn4, zn5 bool
	for i := 1; i <= max; i++ {
		if i <= n1 {
			zn1 = true
		} else {
			zn1 = false
		}
		if i <= n2 {
			zn2 = true
		} else {
			zn2 = false
		}
		if i <= n3 {
			zn3 = true
		} else {
			zn3 = false
		}
		if i <= n4 {
			zn4 = true
		} else {
			zn4 = false
		}
		if i <= n5 {
			zn5 = true
		} else {
			zn5 = false
		}
		line2(zn1, zn2, zn3, zn4, zn5)
	}
}
func main() {
	var n1, n2, n3, n4, n5 int
	fmt.Print("n1: ")
	fmt.Scan(&n1)
	fmt.Println()
	fmt.Print("n2: ")
	fmt.Scan(&n2)
	fmt.Println()
	fmt.Print("n3: ")
	fmt.Scan(&n3)
	fmt.Println()
	fmt.Print("n4: ")
	fmt.Scan(&n4)
	fmt.Println()
	fmt.Print("n5: ")
	fmt.Scan(&n5)
	fmt.Println()
	draw11(n1, n2, n3, n4, n5)
}
