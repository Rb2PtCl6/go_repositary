package main

import "fmt"

func main() {
	fmt.Println("Эта программа вычисляет среднее")
	fmt.Println("арифметическое нескольких ненулевых чисел.")
	fmt.Println()
	var x, sum float64
	var count int
	for {
		fmt.Printf("  Введите %v-е число  /^D - конец ввода/ :  ", count)
		if n, _ := fmt.Scan(&x); n == 0 {
			break
		}
		count++
		sum += x
	}
	fmt.Println()
	fmt.Printf("Введено %v чисел.\n", count)
	if count > 0 {
		fmt.Printf("Их сумма равна  %7.4f,\n", sum)
		fmt.Printf("среднее арифметическое равно  %7.4f.\n", sum/float64(count))
	}
}
