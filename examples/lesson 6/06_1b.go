package main

import (
	"fmt"
)

func prime(n int) bool {
	if n < 2 {
		return 1 == 0
	}
	d := 2
	for n%d != 0 {
		d++
	}
	return n == d
}

func main() {
	fmt.Println("Эта программа находит 10 представлений чётного")
	fmt.Println("натурального числа в виде разности двух простых чисел.")
	fmt.Println()

	var n int

	for (n <= 0) || (n%2 == 1) {
		fmt.Print("Введите чётное натуральное число N: ")
		fmt.Scan(&n)
	}

	for a, count := 3, 1; count <= 10; a, count = a+2, count+1 {
		b := a + n
		for !prime(a) || !prime(b) {
			a, b = a+2, b+2
		}
		fmt.Printf("%4d. %d - %d = %d\n", count, b, a, n)
	}
}
