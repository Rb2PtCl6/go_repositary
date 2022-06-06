package main

import "fmt"

func main() {
	fmt.Println("Эта программа вычисляет среднее")
	fmt.Println("арифметическое нескольких ненулевых чисел.")
	fmt.Println()
	sum := 0.0
	count := 0
	x := 1.0
	for x != 0.0 {
		count++
		fmt.Printf("  Введите %v-е число  /0 - конец ввода/  :  ", count)
		fmt.Scan(&x)
		sum += x
	}
	count--
	fmt.Println()
	fmt.Printf("Введено %v чисел.\n", count)
	if count > 0 {
		fmt.Printf("Их сумма равна  %7.4f,\n", sum)
		fmt.Printf("среднее арифметическое равно  %7.4f.\n", sum/float64(count))
	}
}
