package main

import (
	"fmt"
)

// delete возвращает строку, получающуюся при удалении из строки s
// фрагмента от символа #from и до символа #to включительно
func delete(s string, from, to int) string {
	r := []rune(s)
	if from < 0 {
		from = 0
	}
	to++
	if to > len(r) {
		to = len(r)
	}
	if from < to {
		s = string(r[:from]) + string(r[to:])
	}
	return s
}

func main() {
	fmt.Println(delete("abcdefghij", 5, 7))  //  abcdeij
	fmt.Println(delete("abcdefghij", 7, 2))  //  abcdefghij
	fmt.Println(delete("abcdefghij", 7, 12)) //  abcdefg
	fmt.Println(delete("abcdefghij", 4, 4))  //  abcdfghij
	fmt.Println(delete("abcdefghij", 0, 0))  //  bcdefghij
}
