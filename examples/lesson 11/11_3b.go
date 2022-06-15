package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Считываем строки из файла strings.dat, переворачиваем
	// их и записываем в файл strings.res

	fin, _ := os.Open("strings.dat")
	defer fin.Close()

	fout, _ := os.Create("strings.res")
	defer fout.Close()

	// Create a new Scanner for the file.
	scanner := bufio.NewScanner(fin)
	for scanner.Scan() {
		s := scanner.Text()
		Reverse1(&s)
		fmt.Fprintln(fout, s)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода: ", err)
		os.Exit(1)
	}
}

func Reverse1(ps *string) {
	r := []rune(*ps)
	res := make([]rune, len(r))
	for i, k := 0, len(r); i < len(r); i++ {
		k--
		res[k] = r[i]
	}
	*ps = string(res)
}

func Reverse2(ps *string) {
	res := ""
	for _, x := range *ps {
		res = string(x) + res
	}
	*ps = res
}

func Reverse3(ps *string) {
	r := []rune(*ps)
	for left, right := 0, len(r)-1; left < right; {
		r[left], r[right] = r[right], r[left]
		left++
		right--
	}
	*ps = string(r)
}
