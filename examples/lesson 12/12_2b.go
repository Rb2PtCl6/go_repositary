package main

import (
	"bufio"
	"fmt"
	"os"
)

var abc = []rune("абвгдеёжзийклмнопрстуфхцчшщъыьэюя")

func main() {
	fmt.Printf("Эта программа подсчитывает, сколько")
	fmt.Printf("различных маленьких русских букв Вы ввели.\n\n")
	fmt.Printf("Вводите строки одну за другой. Конец ввода - Ctrl+Z\n")

	counter := make([]uint, len(abc))
	n, in := 0, bufio.NewScanner(os.Stdin)
	for in.Scan() {
		n++
		for _, ch := range in.Text() {
			if i := runeIndex(abc, ch); i >= 0 {
				counter[i]++
			}
		}
	}

	fmt.Println("Всего ввели", n, "строк.")
	for i, cnt := range counter {
		if cnt > 0 {
			fmt.Printf("%c  %4d\n", abc[i], cnt)
		}
	}
}

func runeIndex(s []rune, ch rune) int {
	for i, c := range s {
		if ch == c {
			return i
		}
	}
	return -1
}
