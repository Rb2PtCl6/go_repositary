package main

import (
	"fmt"
)

func main() {
	var speed, time, gone float64
	fmt.Print("speed: ")
	fmt.Scan(&speed)
	fmt.Println()
	fmt.Print("time: ")
	fmt.Scan(&time)
	fmt.Println()
	gone = speed * time
	fmt.Print("gone: ", gone)
}
