package main

import (
	"fmt"
)

func main() {
	var t1, t2, sum, sum2, sum3 float64

	fmt.Print("  Введите км  ")
	fmt.Scan(&t1)
	fmt.Print("Введите часы  ")
	fmt.Scan(&t2)
	sum = t2 / 60
	fmt.Print(t1, "-kilometr", sum, "-min")
	sum2 = t1 / 1000
	sum3 = t2 * 3600
	fmt.Print(sum2, "-metr", sum3, "-sek")
}
