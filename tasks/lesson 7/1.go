package main

import (
	"fmt"
)

type tArNum [7]int

func actions() (sum, min, max, aver, idsq int) {
	num1 := tArNum{}
	idsq = 0
	for i := 0; i < 7; i++ {
		fmt.Printf("Number %v : ", i+1)
		fmt.Scan(&num1[i])
		fmt.Println()
	}
	for i, _ := range num1 {
		sum += num1[i]
		if i == 1 {
			min = num1[i]
			max = num1[i]
		} else {
			if num1[i] > max {
				max = num1[i]
			}
			if num1[i] < min {
				min = num1[i]
			}
		}
		if check1(num1[i]) {
			idsq++
		}
	}
	aver = sum / len(num1)
	return
}
func check1(num int) bool {
	var i int = 1
	var rez int
	for {
		rez = i * i
		if rez == num {
			return true
		} else if rez > num {
			return false
		}
		i++
	}
}
func main() {
	sum, min, max, aver, idsq := actions()
	fmt.Printf("sum: %v\n", sum)
	fmt.Printf("min: %v\n", min)
	fmt.Printf("max: %v\n", max)
	fmt.Printf("aver: %v\n", aver)
	fmt.Printf("idsq: %v\n", idsq)
}
