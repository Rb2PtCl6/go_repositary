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
	for ; in.Scan() && in.Text() != ""; n++ {
		// Extra in.Text() ^^^ possibly makes extra memory allocation
		for _, ch := range in.Text() {
			if ch >= 'a' && ch <= 'z' {
				counter[ch] += 1
			}
		}
	}

	fmt.Println("Всего ввели", n, "строк.")
	for ch := 'a'; ch <= 'z'; ch++ {
		if counter[ch] > 0 {
			fmt.Printf("%c  %4d\n", ch, counter[ch])
		}
	}
}
