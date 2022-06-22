package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type tHockey struct {
	//example: 10,1,1999,Ventspils Masti,4,Ventspils Trubas,6,
	date        tDate  //from example: 10,1,1999
	owner       string //from example: Ventspils Masti
	owner_goals int    //from example: 4
	guess       string //from example: Ventspils Trubas
	guess_goals int    //from example: 6
	isOwertime  bool   //from example:  " "/"*" false/true
}
type tDate struct {
	//example: 10,1,1999
	day    int //from example: 10
	mounth int //from example: 1
	year   int //from example: 1999
}

const file_name string = "hockey.csv"

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
func read_file() (base []tHockey) {
	file, err := os.Open(file_name)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	reader := bufio.NewScanner(file)
	i := 0
	max := max_num()
	base = make([]tHockey, max)
	for reader.Scan() {
		str := strings.Split(reader.Text(), ",")
		base[i].date.day, _ = strconv.Atoi(str[0])
		base[i].date.mounth, _ = strconv.Atoi(str[1])
		base[i].date.year, _ = strconv.Atoi(str[2])
		base[i].owner = str[3]
		base[i].owner_goals, _ = strconv.Atoi(str[4])
		base[i].guess = str[5]
		base[i].guess_goals, _ = strconv.Atoi(str[6])
		if str[7] == "*" {
			base[i].isOwertime = true
		} else {
			base[i].isOwertime = false
		}
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
