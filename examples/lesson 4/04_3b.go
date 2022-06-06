package main

import "fmt"

func main() {

	fmt.Println("Эта программа проверяет, является ли")
	fmt.Println("введенное натуральное число степенью двойки")
	fmt.Println()
	var n int
	for n <= 0 {
		fmt.Print("  Введите натуральное число: ")
		fmt.Scan(&n)
	}

	var degree int
	nn := n
	for nn%2 == 0 {
		nn /= 2     //  nn = nn / 2
		degree += 1 //  degree = degree + 1
	}
	fmt.Println()
	if nn > 1 {
		fmt.Printf("Число %d не является степенью двойки.\n", n)
	} else {
		fmt.Printf("Число %d является %5d-й   степенью двойки.\n", n, degree)
	}
}
