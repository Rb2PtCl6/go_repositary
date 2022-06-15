package main

import (
	"fmt"
	"os"
)

type TDataArray []int

func main() {
	// Считываем числа из файла numbers2.dat
	// Выводим их в обратном порядке в файл numbers2.res

	fin, err := os.Open("numbers2.dat")
	if err != nil {
		return
	}
	defer fin.Close()

	var a TDataArray
	for {
		var c int
		_, err = fmt.Fscanf(fin, "%d\n", &c)
		if err != nil {
			break
		}
		a = append(a, c)
	}

	fout, err := os.Create("numbers2.res")
	if err != nil {
		return
	}
	defer fout.Close()

	for i := len(a) - 1; i >= 0; i-- {
		_, err = fmt.Fprintln(fout, a[i])
		if err != nil {
			return
		}
	}
}
