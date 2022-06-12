package main

import (
	"fmt"
)

// sum5 возвращает сумму элементов пятиэлементного массива
func sum5(c [5]int) int {
	var res int
	for i := 0; i < 5; i++ {
		res += c[i]
	}
	return res
}

func main() {
	a := [5]int{21, 3, 12, 5, 31}
	fmt.Println(sum5(a))
}
