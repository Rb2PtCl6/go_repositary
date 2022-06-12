package main

import "fmt"

func main() {
	fmt.Println("Программа вычисляет такую последовательность:")
	fmt.Println("ее K-й член - это число, образованное двумя")
	fmt.Println("последними цифрами числа N^K.")
	fmt.Println()

	var n int
	for n <= 0 || n >= 100 {
		fmt.Print("Введите N (от 1 до 99): ")
		fmt.Scan(&n)
	}

	x := [51]int{1}
	var k, i int
	for i == k {
		k++
		x[k] = (x[k-1] * n) % 100
		for i = 0; x[i] != x[k]; i++ {
		}
	}
	if i > 0 {
		fmt.Printf("\nПолучили последовательность:\n")
		for j := 0; j < i; j++ {
			fmt.Printf("%4d", x[j])
		}
		fmt.Printf("\nДалее последовательность периодически повторяется:\n")
	} else {
		fmt.Printf("\nПоследовательность периодически повторяется:\n")
	}
	for j := i; j < k; j++ {
		fmt.Printf("%4d", x[j])
	}
	fmt.Printf("\nДлина периода равна %d.\n", k-i)
}
