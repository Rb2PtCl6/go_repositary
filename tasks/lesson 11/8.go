package main

import (
	"bufio"
	"fmt"
	"os"
)

func replace() (string, error) {
	file, err := os.Open(file1)
	if err != nil {
		return "error: no result", err
	}
	defer file.Close()
	pth1 := bufio.NewReader(file)
	tmp := []rune(pth1)
	return str, nil
}
func main() {
	var some float64
	fmt.Print("text: ")
	fmt.Scan(&some)
	fmt.Println()
}
