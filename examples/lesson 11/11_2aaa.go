package main

import (
	"fmt"
)

// delete возвращает строку, получающуюся при удалении из строки s
// фрагмента от символа #from и до символа #to включительно
func delete(s string, from, to int) string {
	if from < 0 {
		from = 0
	}
	to++
	if from < to {
		result := make([]rune, 0, len(s))
		i := 0
		for _, r := range s {
			if i < from || to <= i {
				result = append(result, r)
			}
			i++
		}
		s = string(result)
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
