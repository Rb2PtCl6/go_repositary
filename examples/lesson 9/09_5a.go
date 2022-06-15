package main

import (
	"fmt"
)

func printValue(a int) {
	fmt.Println("value of a in deferred function", a)
}

func main() {
	parameter := 5
	defer printValue(parameter)
	parameter = 10
	fmt.Println("value of a before deferred function call", parameter)
}
