package main

import (
	"fmt"
)

func main() {
	var a [10]int
	//  Ввод данных
	for i := 0; i < 10; i++ {
		fmt.Printf("Enter number # %d: ", i+1)
		fmt.Scan(&a[i])
	}
	//  Вывод введённых чисел в обратном порядке
	for i := 10; i > 1; {
		i--
		fmt.Printf("%d, ", a[i])
	}
	fmt.Println(a[0])
}
