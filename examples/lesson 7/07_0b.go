package main

import (
	"fmt"
)

func main() {
	var a [3]int //   массив из трёх int'ов
	for i := 0; i < 3; i++ {
		fmt.Print(a[i], " ") //0 0 0
	}
}
