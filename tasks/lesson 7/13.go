package main

import (
	"fmt"
)

func main() {
	num := [200]int{0, 1}
	fmt.Println(num[0])
	fmt.Println(num[1])
	for i := 2; i < len(num); i++ {
		num[i] = num[i-1] + num[i-2]
		fmt.Println(num[i])
	}

}
