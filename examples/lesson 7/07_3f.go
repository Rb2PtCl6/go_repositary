package main

import (
	"fmt"
)

const length = 5

type tNumbers [length]int

func main() {
	a := tNumbers{21, 3, 12, 5, 31}
	fmt.Println(sum(a))
}

// sum возвращает сумму элементов массива типа tNumbers
func sum(c tNumbers) int {
	var res int
	for _, x := range c {
		res += x
	}
	return res
}
