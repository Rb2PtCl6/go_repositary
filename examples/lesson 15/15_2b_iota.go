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

func FingerName(finger int) string {
	switch finger {
	case Thumb:  return "Thumb"
	case Index:  return "Index"
	case Middle: return "Middle"
	case Ring: 	 return "Ring"
	case Pinky:  return "Pinky"
	default: return "You are an octopus or an alien..."
	}
}

func main() {
	var finger int
	fmt.Print("Enter finger number: ")
	fmt.Scanln(&finger)
	fmt.Println(FingerName(finger))
}
