package main

import (
	"fmt"
)

const maxLength = 50

type tNumbers [maxLength]int

func main() {
	a := tNumbers{21, 3, 12, 5, 31, 20, 11}
	fmt.Println(sum(a, 5))
}

// sum возвращает сумму первых length элементов массива типа tNumbers
func sum(c tNumbers, length int) int {
	var res int
	for i := 0; i < length; i++ {
		res += c[i]
	}
	return res
}
