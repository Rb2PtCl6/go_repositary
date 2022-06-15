package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Эта программа извлекает из строки все цифры\n\n")

	var str string
	fmt.Println("Введите строку:")
	fmt.Scan(&str)
	str, digits := ExtractDigits1(str)
	digitslen := strLength(digits)
	fmt.Println()
	fmt.Println("Извлечено цифр: ", digitslen)
	if digitslen > 0 {
		fmt.Println("Вот они:", digits)
	}
	fmt.Println()
	fmt.Println("Осталось символов во введенной строке: ", strLength(str))
	fmt.Println(str)
}

// strLength возвращает длину строки s в рунах
func strLength(s string) int {
	return len([]rune(s))
}

func ExtractDigits1(s string) (nodigits string, digits string) {
	r := []rune(s)
	for k := 0; k < len(r); {
		ch := r[k]
		if '0' <= ch && ch <= '9' {
			digits += string(ch)
			r = append(r[:k], r[k+1:]...)
		} else {
			k++
		}
	}
	nodigits = string(r)
	return
}

func ExtractDigits2(s string) (nodigits string, digits string) {
	digitSet := "0123456789"
	for _, ch := range s {
		if strings.ContainsRune(digitSet, ch) {
			digits += string(ch)
		} else {
			nodigits += string(ch)
		}
	}
	return
}
