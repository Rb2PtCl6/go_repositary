package main

import (
	"fmt"
)

func main() {
	a := make([]int, 50)
	a[0] = 0
	a[1] = 1
	for i := 2; i < len(a); i++ {
		a[i] = a[i-1] + a[i-2]
	}
	for i := 0; i < len(a); i++ {
		a[i] = a[i] * a[i]
		fmt.Println(a[i])
	}
}
