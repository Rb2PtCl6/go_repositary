package main

import (
	"fmt"
)

func main() {
	var nm1, nm2, sm float64
	fmt.Print("nm1")
	fmt.Scan(&nm1)
	fmt.Println()
	fmt.Print("nm1")
	fmt.Scan(&nm2)
	fmt.Println()
	sm = nm1 + nm2
	fmt.Println(sm)
	sm = nm1 - nm2
	fmt.Println(sm)
	sm = nm1 * nm2
	fmt.Println(sm)
	sm = nm1 / nm2
}
