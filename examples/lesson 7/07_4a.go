package main

import "fmt"

type tArr [3]int

func f(a tArr) tArr {
	a[0] = 10
	return a
}

func main() {
	a := tArr{21, 8, 93}
	fmt.Println(a) // [21 8 93]
	a = f(a)
	fmt.Println(a) // [10 8 93]
}
