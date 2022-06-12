package main

import "fmt"

func digits(a uint) (n0, n1, n2, n3, n4, n5, n6, n7, n8, n9 int) {
	for {
		switch a % 10 {
		case 0:
			n0++
		case 1:
			n1++
		case 2:
			n2++
		case 3:
			n3++
		case 4:
			n4++
		case 5:
			n5++
		case 6:
			n6++
		case 7:
			n7++
		case 8:
			n8++
		case 9:
			n9++
		}
		a /= 10
		if a == 0 {
			break
		}
	}
	return
}

func main() {
	var a uint
	fmt.Scan(&a)
	fmt.Println(digits(a))
}
