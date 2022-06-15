package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		os.Exit(1)
	}
}

func main() {
	// Считываем числа из файла numbers1.dat
	// Выводим их в обратном порядке в файл numbers1.res

	fin, err := os.Open("numbers1.dat")
	check(err)
	defer fin.Close()

	var n int
	in := bufio.NewReader(fin)
	fmt.Fscanf(in, "%d\n", &n)
	a := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(in, "%d\n", &a[i])
	}

	fout, err := os.Create("numbers1.res")
	check(err)
	defer fout.Close()
	out := bufio.NewWriter(fout)
	defer out.Flush()
	fmt.Fprintln(out, n)
	for i := len(a) - 1; i >= 0; i-- {
		fmt.Fprintln(out, a[i])
	}
}
