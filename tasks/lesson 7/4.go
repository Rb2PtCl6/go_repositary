package main

import (
	"fmt"
)

type tArNum [7]int

func compare() (l0, r0, m0 int) {
	num := tArNum{}
	for i := 0; i < 7; i++ {
		fmt.Printf("Number %v : ", i+1)
		fmt.Scan(&num[i])
	}
	for i, _ := range num {
		if num[i] < 0 {
			l0++
		} else if num[i] > 0 {
			m0++
		} else {
			r0++
		}
	}
	return
}

func main() {
	less, ravn, more := compare()
	fmt.Printf("less 0: %v\n", less)
	fmt.Printf("ravno 0: %v\n", ravn)
	fmt.Printf("more 0: %v\n", more)

}
