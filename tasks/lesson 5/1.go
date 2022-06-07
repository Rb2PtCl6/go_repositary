package main

import (
	"fmt"
)

func check(n int) bool {
	if (n % 2) == 0 {
		return true
	} else {
		return false
	}
}
func main() {
	var num int
	fmt.Print("num: ")
	fmt.Scan(&num)
	if check(num) {
		fmt.Println("Чётное")
	} else {
		fmt.Println("Нечётное")
	}
}
