package main

import (
	"fmt"
	"os"
)

func main() {

	fin, _ := os.Open("numbers.dat")
	defer fin.Close()

	// Чтение данных из файла numbers.dat 
	var a []int
	for {
		var c int
		if _, err := fmt.Fscan(fin, &c); err != nil {
			break
		}
		a = append(a, c)
	}

	// Сортировка простыми вставками
	for k, c := range a {
		var i int
		for i = k; i > 0 && a[i-1] > c; i-- {
			a[i] = a[i-1]
		}
		a[i] = c
	}

	// Запись отсортированных данных в файл numbers.res
	fout, _ := os.Create("numbers.res")
	defer fout.Close()
	for _, c := range a {
		fmt.Fprintln(fout, c)
	}
}
