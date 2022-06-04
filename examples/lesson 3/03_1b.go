package main

import "fmt"

func main() {
	var a, b, min, max float64

	fmt.Println("Эта программа находит максимум и минимум двух чисел.")
	fmt.Println()
	/* Ввод данных */
	fmt.Print("  Введите первое число:  ")
	fmt.Scan(&a)
	fmt.Print("    Введите второе число:  ")
	fmt.Scan(&b)
	/* Вычисления */
	if a < b {
		min, max = a, b
	} else {
		min, max = b, a
	}
	/* Печать результатов */
	fmt.Println()
	fmt.Println(" Максимум =", max)
	fmt.Println(" Минимум =", min)
}
