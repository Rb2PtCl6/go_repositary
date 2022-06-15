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
	fmt.Println("Эта программа извлекает из строки все цифры")
	fmt.Println()
	fmt.Println("Введите строку:")
	str := scanString()
	digits := ExtractDigits(&str)
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

func strLength(s string) int {
	return len([]rune(s))
}

func ExtractDigits(ps *string) string {
	digitSet := "0123456789"
	r := []rune(*ps)
	digits := ""
	for k := 0; k < len(r); {
		if strings.ContainsRune(digitSet, r[k]) {
			digits += string(r[k])
			r = append(r[:k], r[k+1:]...)
		} else {
			k++
		}
	}
	*ps = string(r)
	return digits
}
