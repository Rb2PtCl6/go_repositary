package main

import (
	"fmt"
)

func prime(n int) bool {
	if n < 2 {
		return 0 == 1
	}
	if n%2 == 0 {
		return n == 2
	}
	d := 3
	for (n%d != 0) && (d*d <= n) {
		d += 2
	}
	return (n%d != 0) || (n == 3)
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

	for a, count := 3, 1; count <= 10; a = a + 2 {
		b := a + n
		if prime(a) && prime(b) {
			fmt.Printf("%4d. %d - %d = %d\n", count, b, a, n)
			count++
		}
	}
}
