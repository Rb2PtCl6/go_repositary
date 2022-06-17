package main

import (
	"bufio"
	"fmt"
	"os"
)

const abc = "абвгдеёжзийклмнопрстуфхцчшщъыьэюя"

var m map[rune]bool

func init() {
	m = make(map[rune]bool)
	for _, ch := range abc {
		m[ch] = true
	}
}

func main() {
	fmt.Printf("Эта программа подсчитывает сколько ")
	fmt.Printf("различных пар русских букв Вы ввели.\n\n")
	fmt.Printf("Вводите строки одну за другой. Конец ввода - Ctrl+Z\n")

	n, counter := 0, make(map[string]uint)
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		ch1 := '\n'
		for _, ch2 := range in.Text() {
			if m[ch1] && m[ch2] {
				counter[string(ch1)+string(ch2)]++
				n++
			}
			ch1 = ch2
		}
	}

	fmt.Printf("%4s", "")
	for _, ch2 := range abc {
		fmt.Printf("%4c", ch2)
	}
	fmt.Printf("\n")
	for _, ch1 := range abc {
		fmt.Printf("%4c", ch1)
		for _, ch2 := range abc {
			if np := counter[string(ch1)+string(ch2)]; np == 0 {
				fmt.Printf("%4c", '.')
			} else {
				fmt.Printf("%4d", np)
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("Всего %d пар.\n", n)
}
