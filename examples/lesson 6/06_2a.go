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
	extremeLine(width)
	for row := 1; row < height-1; row++ {
		internalLine(width)
	}
	extremeLine(width)
}

func extremeLine(length int) {
	for i := 0; i < length; i++ {
		fmt.Print("x")
	}
	fmt.Println()
}

func internalLine(length int) {
	fmt.Print("x")
	for i := 1; i < length-1; i++ {
		fmt.Print(" ")
	}
	fmt.Println("x")
}
