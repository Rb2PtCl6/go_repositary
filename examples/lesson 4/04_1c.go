package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Решить 10 примеров на умножение. При первой ошибке программа вылетает.
// Мы хотим знать количество правильных ответов и после выхода из цикла.
func main() {
	rand.Seed(int64(time.Now().Nanosecond()))
	ans := 1
	for ; ans <= 10; ans++ {
		a := rand.Intn(10) + 1
		b := rand.Intn(10) + 1
		fmt.Printf("%d x %d = ", a, b)
		var res int
		fmt.Scan(&res)
		if res != a*b {
			break
		}
	}
	ans--
	fmt.Println(ans, "correct answer(s).")
}
