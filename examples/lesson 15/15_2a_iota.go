package main

import (
	"fmt"
)

const (
	Thumb = iota + 1
	Index
	Middle
	Ring
	Pinky
)	

func main() {
	var finger int
	fmt.Print("Enter finger number: ")
	switch fmt.Scanln(&finger); finger {
	case Thumb:
		fmt.Println("Thumb")
	case Index:
		fmt.Println("Index")
	case Middle:
		fmt.Println("Middle")
	case Ring:
		fmt.Println("Ring")
	case Pinky:
		fmt.Println("Pinky")
	}
}
