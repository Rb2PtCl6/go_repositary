package main

import (
	"fmt"
)

func main() {
	comp1, comp2 := [8]int{}, [8]int{}
	var rez int
	for i := 0; i < len(comp1); i++ {
		fmt.Print("First/second: ")
		fmt.Scan(&comp1[i], &comp2[i])
		fmt.Println()
		if comp1[i] == comp2[i] {
			rez = i
		}
	}
	fmt.Println("rez: ", rez)
}
