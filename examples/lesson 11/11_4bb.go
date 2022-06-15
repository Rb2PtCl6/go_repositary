package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// scanString читает одну строку, включая пробелы
func scanString() string {
	in := bufio.NewReader(os.Stdin)
	str, err := in.ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Println("Ошибка ввода: ", err)
	}
	return str
}

func main() {
	var str, digits string
	fmt.Println("Эта программа извлекает из строки все цифры")
	fmt.Println()
	fmt.Println("Введите строку:")
	str = scanString()
	str, digits = ExtractDigits1(str)
	//digitslen := strLength(digits)
	digitsLen := len(digits)
	fmt.Println()
	fmt.Println("Извлечено цифр: ", digitsLen)
	if digitsLen > 0 {
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

func ExtractDigits1(s string) (str string, digits string) {
	digitSet := "0123456789"
	r := []rune(s)
	for k := 0; k < len(r); {
		ch := r[k]
		if strings.ContainsRune(digitSet, ch) {
			digits += string(ch)
			r = append(r[:k], r[k+1:]...)
		} else {
			k++
		}
	}
	str = string(r)
	return
}

func ExtractDigits2(s string) (nodigits string, digits string) {
	for _, ch := range s {
		if '0' <= ch && ch <= '9' {
			digits += string(ch)
		} else {
			nodigits += string(ch)
		}
	}
	return
}
