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
		if remainder = dividend % divisor; remainder == 0 {
			fmt.Printf("%d делится на %d без остатка.\n", dividend, divisor)
			fmt.Printf("%d : %d = %d\n", dividend, divisor, dividend/divisor)
		} else { // не делится - остаток не 0
			fmt.Printf("%d не делится на %d.\n", dividend, divisor)
			fmt.Printf("%d : %d = %d и %d в остатке.",
				dividend, divisor, dividend/divisor, remainder)
		}
	}
}
