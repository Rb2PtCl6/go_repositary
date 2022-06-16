package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func isRussian(s rune) bool {
	const russian = "абвгдеёжзийклмнопрстуфхцчшщъыьэюя"
	russian1 := []rune(russian)
	for _, rus := range russian1 {
		if s == rus {
			return true
		}
	}
	return false
}
func isSmall(s rune) bool {
	const russian = "абвгдеёжзийклмнопрстуфхцчшщъыьэюя"
	russian1 := []rune(russian)
	for _, rus := range russian1 {
		if s == rus {
			return true
		}
	}
	return false
}
func replace1() string {
	file, err := os.Open("../lesson 5/2.go")
	var rez string
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	file_r := bufio.NewScanner(file)
	for file_r.Scan() {
		file_t := []rune(file_r.Text())
		for i, simb := range file_t {
			if isRussian(simb) {
				if isSmall(simb) {
					file_t[i] = unicode.ToUpper(file_t[i])
				}
			}
		}
		rez += string(file_t)
	}
	return rez
}
func main() {
	fmt.Println(replace1())
}
