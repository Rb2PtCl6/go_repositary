package main

import (
	"fmt"
)

func main() {
	name := "ciąg znaków"
	// name[0] = 'C'        cannot assign to name[0]
	fmt.Println("1. - - - - - - - - -")
	runes := []rune(name)
	for _, c := range runes {
		fmt.Printf("%c", c)
	}
	fmt.Println()
	fmt.Println(name)
	fmt.Println(runes)
	fmt.Println("2. - - - - - - - - -")
	runes[0] = 'C'
	for _, c := range runes {
		fmt.Printf("%c", c)
	}
	fmt.Println()
	fmt.Println(name)
	fmt.Println(runes)
	fmt.Println("3. - - - - - - - - -")
	runes[2] = 'A'
	for _, c := range runes {
		fmt.Printf("%c", c)
	}
	fmt.Println()
	fmt.Println(name)
	fmt.Println(runes)
	name = string(runes)
	fmt.Println()
	fmt.Println(name)
}
