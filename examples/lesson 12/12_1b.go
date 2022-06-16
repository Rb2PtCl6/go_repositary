package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// В файле strings.dat переворачиваем все слова задом наперёд.
	// Предыдущую версию файла сохраняем в strings.bak.

	processing("strings.dat", "strings.tmp")

	// files manipulation
	if exists("strings.bak") {
		os.Remove("strings.bak")
	}
	os.Rename("strings.dat", "strings.bak")
	os.Rename("strings.tmp", "strings.dat")
}

func processing(source string, target string) {

	// Create a new Scanner for the input
	fin, _ := os.Open(source)
	defer fin.Close()
	scanner := bufio.NewScanner(fin)

	// Create a new Writer for the output
	fout, _ := os.Create(target)
	defer fout.Close()
	writer := bufio.NewWriter(fout)
	defer writer.Flush()

	for scanner.Scan() {
		wordlist := strings.SplitAfter(scanner.Text(), " ")
		for _, word := range wordlist {
			fmt.Fprint(writer, reverse(word))
		}
		fmt.Fprintln(writer)
	}
}

func reverse(s string) string {
	r := []rune(s)
	for left, right := 0, len(r)-1; left < right; {
		r[left], r[right] = r[right], r[left]
		left++
		right--
	}
	return string(r)
}

func exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func exists0(name string) bool {
	f, err := os.Open(name)
	if err == nil {
		f.Close()
		return true
	} else {
		return !os.IsNotExist(err)
	}
}
