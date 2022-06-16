package main

import (
	"bufio"
	"fmt"
	"os"
)

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
	map1['а'] = 0
	map1['А'] = 0
	map1['б'] = 0
	map1['Б'] = 0
	map1['и'] = 0
	map1['И'] = 0
	map1['й'] = 0
	map1['Й'] = 0
	for file_r.Scan() {
		file_t := []rune(file_r.Text())
		for _, simb := range file_t {
			if simb == 'а' {
				map1[russian1[0]]++
			} else if simb == 'А' {
				map1[russian1[1]]++
			} else if simb == 'б' {
				map1[russian1[2]]++
			} else if simb == 'Б' {
				map1[russian1[3]]++
			} else if simb == 'и' {
				map1[russian1[4]]++
			} else if simb == 'И' {
				map1[russian1[5]]++
			} else if simb == 'й' {
				map1[russian1[6]]++
			} else if simb == 'Й' {
				map1[russian1[7]]++
			}
		}
	}
	return map1
}
func main() {
	map0 := count1()
	fmt.Printf("а: %d\nА: %d\nб: %d\nБ: %d\nи: %d\nИ: %d\nй: %d\nЙ: %d\n", map0['а'], map0['А'], map0['б'], map0['Б'], map0['и'], map0['И'], map0['й'], map0['Й'])
}
