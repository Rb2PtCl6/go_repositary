package main

import (
	"fmt"
)

func form_arr(width, height int) {
	rez := make([]map[string]int)
}
func main() {
	var width, height int

	fmt.Print("width and height: ")
	fmt.Scan(&width)
	fmt.Println()
	fmt.Print("height: ")
	fmt.Scan(&height)
	fmt.Println()
	form_arr(width, height)
}
