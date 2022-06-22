package main

import (
	"fmt"
)

func getN() (num int) {
	fmt.Scanln(&num)
	return
}

func main() {
	switch num := getN(); { //num is not a constant
	case num < 50:
		fmt.Printf("%d is less than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is less than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is less than 200\n", num)
	default:
		fmt.Printf("%d is greater or equal to 200\n", num)
	}
}
