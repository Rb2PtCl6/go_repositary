package main

import (
	"bufio"
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

	in := bufio.NewReader(fin)
	var a TDataArray
	for {
		var c int
		if _, err = fmt.Fscanf(in, "%d\n", &c); err != nil {
			break
		}
		a = append(a, c)
	}

	fout, err := os.Create("numbers2.res")
	if err != nil {
		return
	}
	defer fout.Close()
	out := bufio.NewWriter(fout)
	defer out.Flush()
	for i := len(a) - 1; i >= 0; i-- {
		if _, err = fmt.Fprintln(out, a[i]); err != nil {
			return
		}
	}
}
