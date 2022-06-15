package main

import (
	"fmt"
	"os"
)

type TDataArray []int

func main() {
	// Считываем числа из файла num.dat
	// Выводим их в обратном порядке в файл num.res

	fin, err := os.Open("num.dat")
	if err != nil {
		os.Exit(1)
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

	fout, err := os.Create("num.res")
	if err != nil {
		os.Exit(1)
	}
	defer fout.Close()

	for i := len(a) - 1; i >= 0; i-- {
		fmt.Fprintf(fout, "%d\n", a[i])
	}
}
