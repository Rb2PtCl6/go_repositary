package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
//В файле snark.txt заменяем все маленькие буквы на большие, а большие - на маленькие.

	var s string

	fin, _ := os.Open("snark.txt")

	fout, _ := os.Create("snark.tmp")

	// Create a new Scanner for the input
	scanner := bufio.NewScanner(fin)

	// Create a new Writer for the output
	writer := bufio.NewWriter(fout)

	for scanner.Scan() {
		s = scanner.Text()
		fmt.Fprintln(writer, changeCase(s))
	}
	writer.Flush()
	fout.Close()
	fin.Close()
	if err := os.Remove("snark.txt"); err == nil  {
		// Кажется, что эта проверка не нужна, os.Rename сам 
		// справится и запишет новый файл поверх старого, но		
		// запросто можно нарваться на какую-нибудь особенность,           
		// например, файл snark.txt read-only или ещё как - и 
		// тогда ничего у наc  не получится
		os.Rename("snark.tmp", "snark.txt")
	}  else  {
		os.Remove("snark.tmp")
		fmt.Println("Не удается изменить файл snark.txt")
	}	
}

func changeCase(s string) (res string) {
	for _, ch := range s {
		onechar := string(ch)
		if strings.ToUpper(onechar) != onechar {
			res += strings.ToUpper(onechar)
		} else if strings.ToLower(onechar) != onechar {
			res += strings.ToLower(onechar)
		} else {
			res += onechar
		}
	}
	return
}

func changeCase1(s string) (res string) {
	for _, ch := range s {
		onechar := string(ch)
		if up := strings.ToUpper(onechar); up != onechar {
			res += up
		} else if low := strings.ToLower(onechar); low != onechar {
			res += low
		} else {
			res += onechar
		}
	}
	return
}

func changeCase2(s string) (res string) {
	for _, ch := range s {
		if unicode.IsLower(ch) {
			res += string(unicode.ToUpper(ch))
		} else if unicode.IsUpper(ch) {
			res += string(unicode.ToLower(ch))
		} else {
			res += string(ch)
		}
	}
	return
}

