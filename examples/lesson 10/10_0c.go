package main

import (
	"fmt"
)

func main() {
	name := [...]string{"ciąg znaków", "տողը"}
	for _, s := range name {
		runes := []rune(s)
		fmt.Println(s, len(s), len(runes))
		for _, c := range s {
			fmt.Printf("%c ", c)
		}
		fmt.Println()
		for i, _ := range runes {
			fmt.Printf("%c ", runes[i])
		}
		fmt.Printf("\n\n")
	}
}
