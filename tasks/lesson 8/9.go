package main

import (
	"fmt"
)

func tvo_sl(n1, n2 []int) []int {
	rez := make([]int, len(n1)+len(n2))
	i := 0
	for ; i < len(n1); i++ {
		rez[i] = n1[i]
	}
	for ; i < len(rez); i++ {
		rez[i] = n2[i-len(n1)]
	}
	return rez
}
func main() {
	n1, n2 := make([]int, 9), make([]int, 5)
	for i := 0; i < len(n1); i++ {
		fmt.Printf("slice 1-%v: ", i+1)
		fmt.Scan(&n1[i])
		fmt.Println()
	}
	for i := 0; i < len(n2); i++ {
		fmt.Printf("slice 2-%v: ", i+1)
		fmt.Scan(&n2[i])
		fmt.Println()
	}
	rez := tvo_sl(n1, n2)
	fmt.Println(rez)
}
