package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var found int
	fmt.Println("Эта программа находит все разложения ")
	fmt.Println("натурального числа N на сумму двух квадратов.")
	fmt.Println()
	n = 0
	for n <= 0 {
		fmt.Print("  Введите натуральное число N:  ")
		fmt.Scan(&n)
	}
	fmt.Println()

	found = 0
	for n1 := 0; 2*n1*n1 <= n; n1 += 1 {
		n2, ok := perfectSquare(n - n1*n1)
		if ok == 1 {
			fmt.Printf("%d = %d² + %d²\n", n, n1, n2)
			found++
		}
	}

	if found == 0 {
		fmt.Println("Разложений нет")
	}
}

/* perfectSquare возвращает √n, 1, если n является точным квадратом,
   иначе возвращает -1, 0.
*/
func perfectSquare(n int) (sqrt int, yes int) {
	if n >= 0 {
		n2 := int(math.Sqrt(float64(n) + 0.1))
		if n2*n2 == n {
			return n2, 1
		}
	}
	return -1, 0
}
