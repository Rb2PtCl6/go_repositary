package main

import (
	"fmt"
)

// Возвращает длину и сумму цифр целого неотрицательного числа n
func numProps(n uint) (length, sum uint) {
	if n == 0 {
		return 1, 0
	}
	for ; n > 0; n /= 10 {
		length++
		sum += n % 10
	}
	return
}

// Возвращает частное и остаток от деления dividend на divisor
func DivMod(dividend, divisor uint) (quotient, remainder uint) {
	return dividend / divisor, dividend % divisor
}

func main() {
	// Выводит все 4-значные числа, которые делятся на сумму своих цифр
	for n := uint(1000); n < 10000; n++ {
		_, digSum := numProps(n)
		if _, remainder := DivMod(n, digSum); remainder == 0 {
			fmt.Println(n)
		}
	}
}
