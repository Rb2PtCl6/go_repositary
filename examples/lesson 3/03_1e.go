package main

import "fmt"

func main() {
	var min, max float64

	fmt.Println("Эта программа находит максимум и минимум двух чисел.")
	fmt.Println()
	/* Ввод данных */
	fmt.Print("  Введите первое число:  ")
	fmt.Scan(&min)
	fmt.Print("    Введите второе число:  ")
	fmt.Scan(&max)
	/* Вычисления */
	if min > max {
		min, max = max, min
	}
	/* Печать результатов */
	fmt.Println()
	fmt.Println(" Максимум =", max)
	fmt.Println(" Минимум =", min)
}
