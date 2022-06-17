package main

import (
	"bufio"
	"fmt"
	"os"
)

const abc = "абвгдеёжзийклмнопрстуфхцчшщъыьэюя"
const N = len(abc)

var m map[rune]int

func init() {
	m = make(map[rune]int)
	i := 0
	for _, ch := range abc {
		m[ch] = i
		i++
	}
}

func main() {
	fmt.Printf("Эта программа подсчитывает сколько")
	fmt.Printf("различных пар русских букв Вы ввели.\n\n")
	fmt.Printf("Вводите строки одну за другой. Конец ввода - Ctrl+Z\n")

	var counter [N][N]uint
	var n uint
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		ch1 := '\n'
		for _, ch2 := range in.Text() {
			if i1, ok := m[ch1]; ok {
				if i2, ok := m[ch2]; ok {
					counter[i1][i2]++
					n++
				}
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
		i1 := m[ch1]
		fmt.Printf("%4c", ch1)
		for _, ch2 := range abc {
			i2 := m[ch2]
			if np := counter[i1][i2]; np == 0 {
				fmt.Printf("%4c", '.')
			} else {
				fmt.Printf("%4d", np)
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("Всего %d пар.\n", n)
}
