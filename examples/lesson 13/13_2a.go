package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Эта программа подсчитывает сколько ")
	fmt.Printf("различных пар английских букв Вы ввели.\n\n")
	fmt.Printf("Вводите строки одну за другой. Конец ввода - Ctrl+Z\n")

	var counter ['z' - 'a' + 1]['z' - 'a' + 1]uint
	var n uint
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		ch1 := '\n'
		for _, ch2 := range in.Text() {
			if 'a' <= ch1 && ch1 <= 'z' && 'a' <= ch2 && ch2 <= 'z' {
				n++
				counter[ch1-'a'][ch2-'a']++
			}
			ch1 = ch2
		}
	}
	fmt.Printf("Всего пар - %d.\n", n)

	fmt.Printf("%s", "   ")
	for ch2 := 'a'; ch2 <= 'z'; ch2++ {
		fmt.Printf("%3c", ch2)
	}
	fmt.Printf("\n")
	n = 0
	for ch1 := 'a'; ch1 <= 'z'; ch1++ {
		fmt.Printf("%3c", ch1)
		for ch2 := 'a'; ch2 <= 'z'; ch2++ {
			if np := counter[ch1-'a'][ch2-'a']; np == 0 {
				fmt.Printf("%3c", '.')
			} else {
				fmt.Printf("%3d", np)
				n++
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("Всего различных пар - %d.\n", n)
}
