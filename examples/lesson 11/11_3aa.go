package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Create a new Scanner for the input
	scanner := bufio.NewScanner(os.Stdin)

	// Create a new Writer for the output
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// Считываем строки, переворачиваем их и выводим
	for scanner.Scan() {
		s := scanner.Text()
		fmt.Fprintln(writer, Reverse(s))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
		os.Exit(1)
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
