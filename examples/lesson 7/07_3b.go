package main

import (
	"fmt"
)

func main() {
	a := tNumbers{21, 3, 12, 5, 31}
	fmt.Println(sum5(a))
}

type tNumbers [5]int

// sum5 возвращает сумму элементов массива типа tNumbers
func sum5(c tNumbers) int {
	var res int
	for i := 0; i < 5; i++ {
		res += c[i]
	}
	return res
}
