package main

import (
	"fmt"
)

func prime(n int) int {
	if n < 2 {
		return 0
	}
	d := 2
	for n%d != 0 {
		d++
	}
	if n == d {
		return 1
	} else {
		return 0
	}
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

	var a = 3
	for count := 1; count <= 10; count++ {
		b := a + n
		for (prime(a) != 1) || (prime(b) != 1) {
			a, b = a+2, b+2
		}
		fmt.Printf("%4d. %d - %d = %d\n", count, b, a, n)
		a += 2
	}
}
