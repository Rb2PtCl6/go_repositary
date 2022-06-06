// 3n+1 пополам

package main

import "fmt"

func main() {
	var n, counter int

	fmt.Print("Введите натуральное число: ")
	fmt.Scan(&n)

	counter = 0
	for n > 1 {
		if n%2 == 0 {
			n /= 2 	//  n = n / 2
		} else {
			n = 3*n + 1
		}
		counter++ 	//  counter += 1   или   counter = counter + 1
		fmt.Println(counter, ". ", n)
	}

	fmt.Println()
	fmt.Println("Всего ", counter, " шагов.")
}
