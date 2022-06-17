package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

type pair struct{ a, b rune }

func main() {
	fmt.Printf("Эта программа подсчитывает сколько ")
	fmt.Printf("различных пар печатных символов Вы ввели.\n\n")
	fmt.Printf("Вводите строки одну за другой. Конец ввода - Ctrl+Z\n")

	counter := make(map[pair]uint)
	in := bufio.NewReader(os.Stdin)
	if ch, _, e := in.ReadRune(); e == nil {
		p := pair{'\n', ch}
		for {
			p.a = p.b
			p.b, _, e = in.ReadRune()
			if e != nil {
				break
			}
			counter[p]++
		}
	}

	set := make(map[rune]bool)
	for p, _ := range counter {
		ch := p.a
		set[ch] = unicode.IsPrint(ch) && ch != ' '
		ch = p.b
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
			if np := counter[pair{ch1, ch2}]; np == 0 {
				fmt.Printf("%4c", '.')
			} else {
				fmt.Printf("%4d", np)
			}
		}
		fmt.Printf("\n")
	}
}
