package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	abc = "абвгдеёжзийклмнопрстуфхцчшщъыьэюя"
	N   = len(abc)
)

func index(ch rune) (n int) {
	for _, r := range abc {
		if r == ch {
			return
		}
		n++
	}
	return
}

func main() {
	fmt.Printf("Эта программа подсчитывает сколько ")
	fmt.Printf("различных пар русских букв Вы ввели.\n\n")
	fmt.Printf("Вводите строки одну за другой. Конец ввода - Ctrl+Z\n")

	var counter [N][N]uint
	var n uint
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		ch1 := '\n'
		i1 := index(ch1)
		for _, ch2 := range in.Text() {
			i2 := index(ch2)
			if i1 < N && i2 < N {
				counter[i1][i2]++
				n++
			}
			ch1, i1 = ch2, i2
		}
	}

	fmt.Printf("%4s", "")
	for _, ch2 := range abc {
		fmt.Printf("%4c", ch2)
	}
	fmt.Printf("\n")
	for i1, ch1 := range []rune(abc) {
		fmt.Printf("%4c", ch1)
		for i2, _ := range []rune(abc) {
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
