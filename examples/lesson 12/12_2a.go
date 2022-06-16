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
	fmt.Println("Вводите строки одну за другой. Конец ввода - Ctrl+Z")

	var counter ['z' - 'a' + 1]uint
	n, in := 0, bufio.NewScanner(os.Stdin)
	for in.Scan() {
		n++
		for _, ch := range in.Text() {
			if ch >= 'a' && ch <= 'z' {
				counter[ch-'a']++
			}
		}
	}

	fmt.Println("Всего ввели", n, "строк.")
	for i, cnt := range counter {
		if cnt > 0 {
			fmt.Printf("%c  %4d\n", 'a'+i, cnt)
		}
	}
}
