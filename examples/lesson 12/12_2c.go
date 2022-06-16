package main

import (
	"bufio"
	"fmt"
	"os"
)

const abc = "абвгдеёжзийклмнопрстуфхцчшщъыьэюя"

var inABC map[rune]bool

func init() {
	inABC = make(map[rune]bool)
	for _, c := range abc {
		inABC[c] = true
	}
}

func main() {
	fmt.Printf("Эта программа подсчитывает, сколько")
	fmt.Printf("различных маленьких русских букв Вы ввели.\n\n")
	fmt.Printf("Вводите строки одну за другой. Конец ввода - Ctrl+Z\n")

	counter := make(map[rune]int)
	n, in := 0, bufio.NewScanner(os.Stdin)
	for in.Scan() {
		n++
		for _, ch := range in.Text() {
			if inABC[ch] {
				counter[ch]++
			}
		}
	}

	fmt.Println("Всего ввели", n, "строк.")
	for _, ch := range abc {
		if counter[ch] > 0 {
			fmt.Printf("%c  %4d\n", ch, counter[ch])
		}
	}
}
