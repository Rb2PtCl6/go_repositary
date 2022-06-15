package main

import (
	"fmt"
)

func main() {
	name := "Rīgas torņi"
	for k, c := range name {
		fmt.Printf("%3d   %c  %6d\n", k, c, c)
	}
}
