package main

import (
	"fmt"
)

// delete() заменяет исходную строку, на строку полученную после удалении
// из строки s фрагмента от символа #from и до символа #to включительно
func delete(ps *string, from int, to int) {
	if from < 0 {
		from = 0
	}
	if len:= strLength(*ps); to >= len  {
		to = len - 1 // номер последнего символа
	}
	if from <= to {
		r := []rune(*ps)
		*ps = string(r[:from]) + string(r[to+1:])
	}
	return
}

// strLength возвращает длину строки s в рунах
func strLength(s string) int {
	return len([]rune(s))
}

func main() {
	var s string
	s = "abcdefghij"
	delete(&s, 5, 7)
	fmt.Println(s)
	s = "abcdefghij"
	delete(&s, 7, 2)
	fmt.Println(s)
	s = "abcdefghij"
	delete(&s, 7, 12)
	fmt.Println(s)
	s = "abcdefghij"
	delete(&s, 4, 4)
	fmt.Println(s)
	s = "abcdefghij"
	delete(&s, 0, 0)
	fmt.Println(s)
}
