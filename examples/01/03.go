package main

import "fmt"

func main() {
	var a, b, diffy, multy float64

	fmt.Println("Эта программа проверяет тождество")
	fmt.Println("   (A-B)(A+B) = A² - B² ")
	fmt.Println()
	fmt.Print("  Введите число A:  ")
	fmt.Scan(&a)
	fmt.Print("  Введите число B:  ")
	fmt.Scan(&b)
	diffy = a*a - b*b
	multy = (a - b) * (a + b)
	fmt.Printf("%5.3f² - %5.3f² = %5.3f", a, b, diffy)
	fmt.Println()
	fmt.Printf("(%5.3f - %5.3f)(%5.3f + %5.3f) = %5.3f", a, b, a, b, multy)
	fmt.Println()
}
