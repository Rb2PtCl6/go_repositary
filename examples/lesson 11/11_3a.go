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

	// Create a new Scanner for the input
	scanner := bufio.NewScanner(fin)

	// Create a new Writer for the output
	writer := bufio.NewWriter(fout)
	defer writer.Flush()

	for scanner.Scan() {
		s := scanner.Text()
		fmt.Fprintln(writer, Reverse(s))
	}
}

func Reverse(s string) string {
	r := []rune(s)
	for left, right := 0, len(r)-1; left < right; {
		r[left], r[right] = r[right], r[left]
		left++
		right--
	}
	return string(r)
}
