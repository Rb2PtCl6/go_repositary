package main

import "fmt"

func main() {
	var dividend, divisor int
	var remainder int

	fmt.Println("Эта программа проверяет делимость двух целых чисел.")
	fmt.Println()
	// Ввод данных
	fmt.Print("  Введите делимое:  ")
	fmt.Scan(&dividend)
	fmt.Print("    Введите делитель:  ")
	fmt.Scan(&divisor)
	fmt.Println()
	// Обработка и вывод результатов
	if divisor == 0 {
		fmt.Println(" На 0 делить нельзя :( ")
	} else {
		remainder = dividend % divisor
		if remainder == 0 {
			fmt.Println(dividend, "делится на", divisor, "без остатка.")
			fmt.Println(dividend, ":", divisor, "=", dividend/divisor)
		} else { // не делится - остаток не 0
			fmt.Println(dividend, "не делится на", divisor)
			fmt.Println(dividend, ":", divisor, "=", dividend/divisor,
				"и", remainder, "в остатке")
		}
	}
}
