package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Эта программа подсчитывает, сколько различных")
	fmt.Println("маленьких английских букв Вы ввели.")
	fmt.Println()
	fmt.Println("Вводите строки одну за другой. Конец ввода - пустая строка")

	counter := make(map[rune]int)
	n, in := 0, bufio.NewScanner(os.Stdin)
	// for in.Scan() && in.Text() != "" { // Extra in.Text() possibly makes extra memory allocation
	for in.Scan() {
		n++
		for _, ch := range in.Text() {
			if ch >= 'a' && ch <= 'z' {
				counter[ch]++
			}
		}
	}

	fmt.Println("Всего введенных строк:", n)
	for ch, count := range counter {
		fmt.Printf("%c  %4d\n", ch, count)
	}
}
