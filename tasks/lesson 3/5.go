package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, dis float64
	fmt.Print("a: ")
	fmt.Scan(&a)
	fmt.Println()
	fmt.Print("b: ")
	fmt.Scan(&b)
	fmt.Println()
	fmt.Print("c: ")
	fmt.Scan(&c)
	fmt.Println()
	dis = b*b - 4*a*c
	if dis < 0 {
		fmt.Println("Корней нет!")
	} else if dis == 0 {
		fmt.Println(-b / 2 * a)
	} else if dis > 0 {
		fmt.Println((-b + math.Sqrt(dis)) / (2 * a))
		fmt.Println((-b - math.Sqrt(dis)) / (2 * a))
	} else {
		fmt.Println("Unknown error!")
	}
}
