package main

import (
	"fmt"
)

const length = 5

type tNumbers [length]int

func main() {
	a := tNumbers{21, 3, 12, 5, 31}
	fmt.Println(sum5(a))
}

// sum5 возвращает сумму элементов массива типа tNumbers
func sum5(c tNumbers) int {
	var res int
	for i := 0; i < len(c); i++ {
		res += c[i]
	}
	return res
}
