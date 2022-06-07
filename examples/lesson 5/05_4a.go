package main

import (
	"fmt"
)

// Возвращает длину и сумму цифр целого неотрицательного числа n
func numProps(n uint) (uint, uint) {
	if n == 0 {
		return 1, 0
	}
	var sum, length uint
	for ; n > 0; n /= 10 {
		length++
		sum += n % 10
	}
	return length, sum
}

func main() {
	var n uint
	fmt.Scan(&n)
	length, digSum := numProps(n)
	fmt.Printf("Число %d состоит из %d цифр; их сумма равна %d\n", n, length, digSum)
}
