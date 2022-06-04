// Пример 03

package main

import "fmt"

func main() {
	x := 15
	x = x + 1
	// x is 16
	x++
	// x is 17
	x += 10 // x = x + 10
	// x is 27
	x-- // x -= 1
	// x is 26
	x++
	x *= 4
	x = (x + 1) / 2
	x *= 2
	x /= 3

	fmt.Println(x)

	x %= 5

	fmt.Println(x) // x = ?
}
