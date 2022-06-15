package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Считываем числа из файла numbers.dat
	fin, _ := os.Open("numbers.dat")
	defer fin.Close()

	in := bufio.NewReader(fin)
	var sum int
	if _, err := fmt.Fscanf(in, "%d\n", &sum); err != nil {
		fmt.Println("Файл пустой.")
		return
	}

	amt := 1
	maximal, nMax := sum, 1
	minimal, nMin := sum, 1
	for {
		var x int
		if _, err := fmt.Fscanf(in, "%d\n", &x); err != nil {
			break
		}
		amt++
		sum += x

		if x > maximal {
			maximal, nMax = x, 1
		} else if x == maximal {
			nMax++
		}

		if x < minimal {
			minimal, nMin = x, 1
		} else if x == minimal {
			nMin++
		}
	}
	fmt.Printf("Всего в файле %d чисел\n", amt)
	fmt.Printf("Сумма чисел = %d\n", sum)
	fmt.Printf("Среднее арифметическое = %7.4f\n", float64(sum)/float64(amt))
	fmt.Printf("Наибольшее число %d встречается %d раз\n", maximal, nMax)
	fmt.Printf("Наименьшее число %d встречается %d раз\n", minimal, nMin)
}
