package main

import (
	"fmt"
)

func main() {
	var width, height int
	fmt.Print("Ширина: ")
	fmt.Scan(&width)
	fmt.Print("Высота: ")
	fmt.Scan(&height)
	extremeLine(width, 'x')
	for row := 1; row < height-1; row++ {
		internalLine(width, 'x', '+')
	}
	extremeLine(width, 'x')
}

func extremeLine(length int, borderSymbol rune) {
	for i := 0; i < length; i++ {
		fmt.Printf("%c", borderSymbol)
	}
	fmt.Println()
}

func internalLine(length int, borderSymbol, fillSymbol rune) {
	fmt.Printf("%c", borderSymbol)
	for i := 1; i < length-1; i++ {
		fmt.Printf("%c", fillSymbol)
	}
	fmt.Printf("%c\n", borderSymbol)
}
