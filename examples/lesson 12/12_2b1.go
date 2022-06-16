package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const abc = "абвгдеёжзийклмнопрстуфхцчшщъыьэюя"

func main() {
	fmt.Printf("Эта программа подсчитывает, сколько")
	fmt.Printf("различных маленьких русских букв Вы ввели.\n\n")
	fmt.Printf("Вводите строки одну за другой. Конец ввода - Ctrl+Z\n")

	counter := make(map[rune]int)
	n, in := 0, bufio.NewScanner(os.Stdin)
	for in.Scan() {
		n++
		for _, ch := range in.Text() {
			if strings.ContainsRune(abc, ch) {
				counter[ch]++
			}
		}
	}

	fmt.Println("Всего ввели", n, "строк.")
	for ch, cnt := range counter {
		fmt.Printf("%c  %4d\n", ch, cnt)
	}
}
