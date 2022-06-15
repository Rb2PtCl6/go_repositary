package main

import (
	"fmt"
	"os"
)

func main() {

	fin, _ := os.Open("numbers.dat")
	defer fin.Close()

	// Чтение данных из файла numbers.dat с одновременной сортировкой
	var a []int
	for {
		var c int
		if _, err := fmt.Fscan(fin, &c); err != nil {
			break
		}
		a = append(a, 0)
		var i int
		for i = len(a) - 2; i >= 0 && a[i] > c; i-- {
			a[i+1] = a[i]
		}
		a[i+1] = c
	}

	// Запись отсортированных данных в файл numbers.res
	fout, _ := os.Create("numbers.res")
	defer fout.Close()
	for _, c := range a {
		fmt.Fprintln(fout, c)
	}
}
