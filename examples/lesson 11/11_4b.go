package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// scanString читает одну строку, включая пробелы
func scanString() string {
	in := bufio.NewReader(os.Stdin)
	str, err := in.ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Println("Ошибка ввода: ", err)
	}
	// Удаляем два последних символа - символы перевода строки
	r:= []byte(str)
	str = string(r[:len(r)-2])
	return str
}

func main() {
	var str, digits string
	fmt.Println("Эта программа извлекает из строки все цифры")
	fmt.Println()
	fmt.Println("Введите строку:")
	str = scanString()
	str, digits = ExtractDigits(str)
	digitsLen := strLength(digits)
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

func ExtractDigits(s string) (nodigits string, digits string) {
	for _, ch := range []rune(s) {
		if '0' <= ch && ch <= '9' {
			digits += string(ch)
		} else {
			nodigits += string(ch)
		}
	}
	return
}
