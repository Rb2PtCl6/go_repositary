package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type tStudent struct {
	//example: Lynch,Leo,10A
	name    string //from example: Leo
	surname string //from example: Lynch
	class   string //from example: 10A
}

const file_name string = "students.csv"

func max_num() int {
	file1, err := os.Open(file_name)
	if err != nil {
		fmt.Println(err)
	}
	defer file1.Close()
	reader := bufio.NewScanner(file1)
	i := 0
	for reader.Scan() {
		if reader.Text() != "" {
			i++
		}
	}
	return i
}
func read_file() (base []tStudent) {
	file, err := os.Open(file_name)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	reader := bufio.NewScanner(file)
	i := 0
	max := max_num()
	base = make([]tStudent, max)
	for reader.Scan() {
		str := strings.Split(reader.Text(), ",")
		base[i].surname = str[0]
		base[i].name = str[1]
		base[i].class = str[2]
		if i == max-2 {
			break
		}
		i++
	}
	return
}
func main() {
	//var some float64
	//fmt.Print("text: ")
	//fmt.Scan(&some)
	base := read_file()
	fmt.Println(base[10])
}
