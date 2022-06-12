package main

import "fmt"

type tArr [10]int

func fill(d *tArr) {
	for i, _ := range d {
		d[i] = i
	}
}

func main() {
	var a tArr
	fmt.Println(a) //	[0 0 0 0 0 0 0 0 0 0]
	fill(&a)
	fmt.Println(a) //	[0 1 2 3 4 5 6 7 8 9]
}
