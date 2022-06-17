package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	fmt.Printf("Эта программа подсчитывает сколько ")
	fmt.Printf("различных пар печатных символов Вы ввели.\n\n")
	fmt.Printf("Вводите строки одну за другой. Конец ввода - Ctrl+Z\n")

	counter := make(map[[2]rune]uint)
	in := bufio.NewReader(os.Stdin)
	if ch, _, e := in.ReadRune(); e == nil {
		pair := [2]rune{'\n', ch}
		for {
			pair[0] = pair[1]
			pair[1], _, e = in.ReadRune()
			if e != nil {
				break
			}
			counter[pair]++
		}
	}

	set := make(map[rune]bool)
	for pair, _ := range counter {
		ch := pair[0]
		set[ch] = unicode.IsPrint(ch) && ch != ' '
		ch = pair[1]
		set[ch] = unicode.IsPrint(ch) && ch != ' '
	}

	abc := make([]rune, 0, len(set))
	for ch, yes := range set {
		if yes {
			abc = append(abc, ch)
		}
	}

	fmt.Printf("%4s", " ")
	for _, ch2 := range abc {
		fmt.Printf("%4c", ch2)
	}
	fmt.Printf("\n")
	for _, ch1 := range abc {
		fmt.Printf("%4c", ch1)
		for _, ch2 := range abc {
			if np := counter[[2]rune{ch1, ch2}]; np == 0 {
				fmt.Printf("%4c", '.')
			} else {
				fmt.Printf("%4d", np)
			}
		}
		fmt.Printf("\n")
	}
}
