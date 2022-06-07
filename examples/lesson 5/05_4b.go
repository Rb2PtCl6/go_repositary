package main

import (
	"fmt"
	"math"
)

// Возвращает длину и сумму цифр целого неотрицательного числа n,
// а также его квадратный корень
func numProps(n uint) (length, sum uint, root2 float64) {
	if n == 0 {
		return 1, 0, 0
	}
	root2 = math.Sqrt(float64(n))
	for ; n > 0; n /= 10 {
		length++
		sum += n % 10
	}
	return
}

func main() {
	var n uint
	fmt.Scan(&n)
	length, digSum, sqroot := numProps(n)
	fmt.Printf("Число %d состоит из %d цифр; их сумма равна %d\n", n, length, digSum)
	fmt.Printf("√%d = %g\n", n, sqroot)
}
