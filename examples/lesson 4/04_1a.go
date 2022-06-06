package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Решить 10 примеров на умножение.
// При первой ошибке программа вылетает.

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 1; i <= 10; i++ {
		a := rand.Intn(10) + 1
		b := rand.Intn(10) + 1
		fmt.Printf("%d x %d = ", a, b)
		var res int
		fmt.Scan(&res)
		if res != a*b {
			fmt.Printf("Error on exercise #%d", i)
			break
		}
	}
}
