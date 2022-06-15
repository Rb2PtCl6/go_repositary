package main

import (
	"fmt"
)

// delete возвращает строку, получающуюся при удалении из строки s
// фрагмента от символа #from и до символа #to включительно
func delete(s string, from int, to int) string {
	if from < 0 {
		from = 0
	}
	if len:= strLength(s); to >= len  {
		to = len - 1 // номер последнего символа
	}
	if from > to {
		return s
	} else {
		r := []rune(s)
		return string(r[:from]) + string(r[to+1:])
	}
}

// strLength возвращает длину строки s в рунах
func strLength(s string) int {
	return len([]rune(s))
}

func main() {
	fmt.Println(delete("abcdefghij", 5, 7))  //  abcdeij
	fmt.Println(delete("abcdefghij", 7, 2))  //  abcdefghij
	fmt.Println(delete("abcdefghij", 7, 12)) //  abcdefg
	fmt.Println(delete("abcdefghij", 4, 4))  //  abcdfghij
	fmt.Println(delete("abcdefghij", 0, 0))  //  bcdefghij
}
