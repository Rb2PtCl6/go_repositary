package main

import (
	"fmt"
	"time"
)

func main() {
	countdown(10)
	fmt.Println("Go!")
}

func countdown(seconds int) {
	for ; seconds > 0; seconds-- {
		fmt.Println(seconds)
		time.Sleep(1 * time.Second)
	}
}
