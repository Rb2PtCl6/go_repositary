package main

import (
	"fmt"
)

// Ищем все натуральные числа k, k < 64, для которых 2^k + 1 делится на k
func main() {
	var n uint64 = 2
	for k := uint64(1); k < 64; k++ {
		if (n+1)%k == 0 {
			fmt.Println(n, k)
		}
		n *= 2
	}
	fmt.Println(n)
}
