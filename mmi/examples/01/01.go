package main

import (
	"fmt"
)

func main() {
	var term1, term2, sum float64

	fmt.Println("Эта программа вычисляет сумму двух чисел")
	fmt.Println()

	fmt.Print("  Введите первое слагаемое:  ")
	fmt.Scan(&term1)
	fmt.Print("    Введите второе слагаемое:  ")
	fmt.Scan(&term2)

	sum = term1 + term2
	fmt.Println()
	fmt.Println(term1, "+", term2, "=", sum)
}
