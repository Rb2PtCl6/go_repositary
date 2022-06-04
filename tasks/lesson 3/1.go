package main

import (
	"fmt"
)

func main() {
	var a, b float64
	var pr string
	fmt.Print("a: ")
	fmt.Scan(&a)
	fmt.Println()
	fmt.Print("b: ")
	fmt.Scan(&b)
	fmt.Println()
	if (a>b){pr="1"} else if (a<b){pr="2"} else if (a==b){pr="0"}else{pr="Unknown error!"}
	fmt.Println(pr)
}