package main

import (
	"bufio"
	"fmt"
	"github.com/nsf/termbox-go"
	"os"
	"sort"
	"strconv"
)

type storage struct {
	amount int
	name   string
	num    int
}

func reader() []storage {
	file, err := os.Open("product.txt")
	if err != nil {
		fmt.Println("Something went wrong!")
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	data := make([]storage, 12)
	mounths := [...]string{
		"январь",
		"февраль",
		"март",
		"апрель",
		"май",
		"июнь",
		"июль",
		"август",
		"сентябрь",
		"октябрь",
		"ноябрь",
		"декабрь",
	}
	i := 0
	for scanner.Scan() {
		data[i].amount, _ = strconv.Atoi(scanner.Text())
		data[i].name = mounths[i]
		data[i].num = i + 1
		i++
	}
	return data
}

func task2() {
	//data := reader()
}
func task3() {
	data := reader()
	sort.SliceStable(data, func(i, j int) bool { return data[i].amount > data[j].amount })
	fmt.Println()
	fmt.Println("mounth numbers")
	fmt.Println()
	for i := 0; i < 12; i++ {
		fmt.Printf("%v ", data[i].num)
	}
	fmt.Println()
	fmt.Println()
	fmt.Println("mount names")
	fmt.Println()
	for i := 0; i < 12; i++ {
		fmt.Printf("%v ", data[i].name)
	}
	fmt.Println()
	fmt.Println()
}
func task4() {
	fmt.Println()
	data := reader()
	//screen
	fmt.Println()
	fmt.Println("Stars")
	fmt.Println()
	for i := 0; i < 12; i++ {
		fmt.Printf("%v ", data[i].amount)
		for i1 := 0; i1 < data[i].amount; i1++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	//txt
	file4, err := os.Create("tmp/barchart1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file4.Close()
	writer := bufio.NewWriter(file4)
	defer writer.Flush()
	fmt.Fprintln(file4)
	fmt.Fprintln(file4, "File")
	fmt.Fprintln(file4)
	for i := 0; i < 12; i++ {
		fmt.Fprintf(file4, "%v ", data[i].amount)
		for i1 := 0; i1 < data[i].amount; i1++ {
			fmt.Fprint(file4, "*")
		}
		fmt.Fprintln(file4)
	}
	//letters
	fmt.Println()
	fmt.Println("Letters")
	fmt.Println()
	letters := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L"}
	for i := 0; i < 12; i++ {
		fmt.Printf("%v ", data[i].amount)
		for i1 := 0; i1 < data[i].amount; i1++ {
			fmt.Print(letters[i])
		}
		fmt.Println()
	}
	println()
	fmt.Println()
}

func task5() {
	data := reader()
	basepos := 0
	err_term := termbox.Init()
	if err_term != nil {
		return
	}
	defer termbox.Close()
	w, h := termbox.Size()
	for i2 := 0; i2 < 12; i2++ {
		if data[i2].amount < 10 {
			termbox.SetCell(0, basepos, data[i2].amount)
		}
	}
}
func task6() {
	//data := reader()
}
func main() {
	//var some float64
	//fmt.Print("text: ")
	//fmt.Scan(&some)
	//fmt.Println()
	if false {
		task2()
	}
	if false {
		task3()
	}
	if false {
		task4()
	}
	if true {
		task5()
	}
	if false {
		task6()
	}

}
