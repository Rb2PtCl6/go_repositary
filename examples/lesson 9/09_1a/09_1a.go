package main

import (
	"fmt"
	"os"
)

func main() {
	// Считываем числа из файла numbers1.dat
	// Выводим их в обратном порядке в файл numbers1.res

	fin, _ := os.Open("numbers1.dat")

	var n int // количество чисел
	fmt.Fscan(fin, &n)

	a := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(fin, &a[i])
	}
	fin.Close()

	fout, _ := os.Create("numbers1.res")

	fmt.Fprintln(fout, n)
	for i := len(a) - 1; i >= 0; i-- {
		fmt.Fprintln(fout, a[i])
	}
	fout.Close()
}
