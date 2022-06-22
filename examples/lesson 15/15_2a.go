package main

import (
	"fmt"
)


func main() {
	var finger int
	fmt.Print("Enter finger number: ")
	fmt.Scanln(&finger)
	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	}
}
