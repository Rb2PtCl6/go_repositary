package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ans := rand.Intn(101)
	for {
		var guess int
		fmt.Print("guess: ")
		fmt.Scan(&guess)
		fmt.Println()
		if guess < ans {
			fmt.Println("Smaller!")
		} else if guess > ans {
			fmt.Println("Bigger!")
		} else {
			fmt.Println("Correct!")
		}
	}

}
