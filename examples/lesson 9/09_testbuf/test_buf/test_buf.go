package main

import (
	"bufio"
	"fmt"
	"os"
)

type TDataArray []int

func check(e error) {
	if e != nil {
		os.Exit(1)
	}
}

func main() {
	// Считываем числа из файла num.dat
	// Выводим их в обратном порядке в файл num.res

	fin, err := os.Open("num.dat")
	check(err)
	defer fin.Close()

	in := bufio.NewReader(fin)
	var a TDataArray
	for {
		var c int
		_, err = fmt.Fscan(in, &c)
		if err != nil {
			break
		}
		a = append(a, c)
	}

	fout, _ := os.Create("num.res")
	defer fout.Close()
	out := bufio.NewWriter(fout)
	defer out.Flush()
	for i := len(a) - 1; i >= 0; i-- {
		_, err = fmt.Fprintln(out, a[i])
		check(err)
	}
}
