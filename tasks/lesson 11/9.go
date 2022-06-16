package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	//"strconv"
)

func scan() string {
	read := bufio.NewReader(os.Stdin)
	str, err := read.ReadString('\n')
	if err != nil && err != io.EOF {
		return "Something went wrong!"
	}
	rez1 := []byte(str)
	str = string(rez1[:len(rez1)-2])
	return str
}
func test() bool {
	str := scan()
	for _, sumb := range str {
		if (sumb < 65 && sumb > 90) || (sumb < 97 && sumb > 122) {
			//A-Z a-z http://yuschikev.narod.ru/Teoria/Inform/kod_tex.html
		}
	}
}
func main() {
	var some float64
	fmt.Print("text: ")
	fmt.Scan(&some)
	fmt.Println()
}
