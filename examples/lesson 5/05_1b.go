package main

import (
	"fmt"
	"math"
)

func sqr(x int) int {
	return x * x
}

// perfectSquare возвращает √n, когда n является точным квадратом.
// Во всех остальных случаях -1.
func perfectSquare(n int) int {
	if n < 0 {
		return -1
	}
	n2 := int(math.Sqrt(float64(n) + 0.1))
	if sqr(n2) == n {
		return n2
	} else {
		return -1
	}
}

func main() {
	fmt.Println("Эта программа находит все разложения")
	fmt.Println("натурального числа N на сумму двух квадратов.")
	fmt.Println()

	var n int
	for n = 0; n <= 0; {
		fmt.Print("Введите N: ")
		fmt.Scan(&n)
	}
	fmt.Println()

	count := 0
	for n1 := 0; 2*sqr(n1) <= n; n1++ {
		var n2 int = perfectSquare(n - sqr(n1))
		if n2 >= 0 {
			count += 1
			fmt.Printf("%4d. %d = %d² + %d²\n", count, n, n1, n2)
		}
	}

	if count == 0 {
		fmt.Println("Разложений нет")
	}
}
