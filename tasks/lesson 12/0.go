package main

import (
	"bufio"
	"fmt"
	"os"
)

func isNesseryLetter(s rune) bool {
	const russian = "аАбБиИйЙ"
	russian1 := []rune(russian)
	for _, rus := range russian1 {
		if s == rus {
			return true
		}
	}
	return false
}
func count1() map[rune]int {
	file, err := os.Open("../lesson 5/2.go")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	file_r := bufio.NewScanner(file)
	const russian = "аАбБиИйЙ"
	russian1 := []rune(russian)
	map1 := make(map[rune]int)
	map1[russian1[0]] = 0
	map1[russian1[1]] = 0
	map1[russian1[2]] = 0
	map1[russian1[3]] = 0
	map1[russian1[4]] = 0
	map1[russian1[5]] = 0
	map1[russian1[6]] = 0
	map1[russian1[7]] = 0
	for file_r.Scan() {
		file_t := []rune(file_r.Text())
		for _, simb := range file_t {
			if simb == russian1[0] {
				map1[russian1[0]]++
			} else if simb == russian1[1] {
				map1[russian1[1]]++
			} else if simb == russian1[2] {
				map1[russian1[2]]++
			} else if simb == russian1[3] {
				map1[russian1[3]]++
			} else if simb == russian1[4] {
				map1[russian1[4]]++
			} else if simb == russian1[5] {
				map1[russian1[5]]++
			} else if simb == russian1[6] {
				map1[russian1[6]]++
			} else if simb == russian1[7] {
				map1[russian1[7]]++
			}
		}
	}
	return map1
}
func main() {
	fmt.Println(count1())
}
