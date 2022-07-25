package main

import (
	"bufio"
	"fmt"
	"time"

	"os"
	"sort"
	"strconv"

	"github.com/nsf/termbox-go"
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
func createpattern(direction string) (pattern [12][20]bool) {
	// "up" or "down"
	data := reader()
	if direction == "down" {
		for i := 0; i < 12; i++ {
			for i3 := 0; i3 < data[i].amount; i3++ {
				pattern[i][i3] = true
			}
			for i4 := data[i].amount; i4 < 20; i4++ {
				pattern[i][i4] = false
			}
		}
	} else {
		for i := 0; i < 12; i++ {
			for i3 := 0; i3 < 20-data[i].amount; i3++ {
				pattern[i][i3] = false
			}
			for i4 := 20 - data[i].amount; i4 < 20; i4++ {
				pattern[i][i4] = true
			}
		}
	}
	return pattern
}
func drawpattern(termboxstatus, isfile, pos int, pattern [12][20]bool) {
	data := reader()
	var x0 int
	if pos == 1 {
		x0 = 1
	}
	if termboxstatus == 0 && isfile == 0 {
		for i := 0; i < 21; i++ {
			if i == 0 {
				for i1 := 0; i1 < 12; i1++ {
					drawnum(0, 0, data[i1].amount, 0, 0)
					fmt.Print(" ")
				}
				fmt.Println()
				continue
			}
			for i2 := 0; i2 < 12; i2++ {
				drawline(0, 0, pattern[i2][i-1], 0, 0)
				fmt.Print(" ")
			}
			fmt.Println()
		}
	} else if termboxstatus == 0 && isfile == 1 {
		file, err := os.Create("tmp/barchart2.txt")
		if err != nil {
			return
		}
		if pos == 2 {
			defer os.Rename("tmp/barchart2.txt", "tmp/barchart3.txt")
		}
		defer file.Close()
		row := bufio.NewWriter(file)
		defer row.Flush()
		for i := 0; i < 21; i++ {
			if pos == 1 && i == 0 {
				for i2 := 0; i2 < 12; i2++ {
					switch data[i2].amount {
					case 0:
						fmt.Fprint(file, "0   ")
						break
					case 1:
						fmt.Fprint(file, "1   ")
						break
					case 2:
						fmt.Fprint(file, "2   ")
						break
					case 3:
						fmt.Fprint(file, "3   ")
						break
					case 4:
						fmt.Fprint(file, "4   ")
						break
					case 5:
						fmt.Fprint(file, "5   ")
						break
					case 6:
						fmt.Fprint(file, "6   ")
						break
					case 7:
						fmt.Fprint(file, "7   ")
						break
					case 8:
						fmt.Fprint(file, "8   ")
						break
					case 9:
						fmt.Fprint(file, "9   ")
						break
					case 10:
						fmt.Fprint(file, "10  ")
						break
					case 11:
						fmt.Fprint(file, "11  ")
						break
					case 12:
						fmt.Fprint(file, "12  ")
						break
					case 13:
						fmt.Fprint(file, "13  ")
						break
					case 14:
						fmt.Fprint(file, "14  ")
						break
					case 15:
						fmt.Fprint(file, "15  ")
						break
					case 16:
						fmt.Fprint(file, "16  ")
						break
					case 17:
						fmt.Fprint(file, "17  ")
						break
					case 18:
						fmt.Fprint(file, "18  ")
						break
					case 19:
						fmt.Fprint(file, "19  ")
						break
					case 20:
						fmt.Fprint(file, "20  ")
						break
					}
				}
				fmt.Fprintln(file)
				continue
			} else if pos == 2 && i == 20 {
				for i2 := 0; i2 < 12; i2++ {
					switch data[i2].amount {
					case 0:
						fmt.Fprint(file, "0   ")
						break
					case 1:
						fmt.Fprint(file, "1   ")
						break
					case 2:
						fmt.Fprint(file, "2   ")
						break
					case 3:
						fmt.Fprint(file, "3   ")
						break
					case 4:
						fmt.Fprint(file, "4   ")
						break
					case 5:
						fmt.Fprint(file, "5   ")
						break
					case 6:
						fmt.Fprint(file, "6   ")
						break
					case 7:
						fmt.Fprint(file, "7   ")
						break
					case 8:
						fmt.Fprint(file, "8   ")
						break
					case 9:
						fmt.Fprint(file, "9   ")
						break
					case 10:
						fmt.Fprint(file, "10  ")
						break
					case 11:
						fmt.Fprint(file, "11  ")
						break
					case 12:
						fmt.Fprint(file, "12  ")
						break
					case 13:
						fmt.Fprint(file, "13  ")
						break
					case 14:
						fmt.Fprint(file, "14  ")
						break
					case 15:
						fmt.Fprint(file, "15  ")
						break
					case 16:
						fmt.Fprint(file, "16  ")
						break
					case 17:
						fmt.Fprint(file, "17  ")
						break
					case 18:
						fmt.Fprint(file, "18  ")
						break
					case 19:
						fmt.Fprint(file, "19  ")
						break
					case 20:
						fmt.Fprint(file, "20  ")
						break
					}
				}
				fmt.Fprintln(file)
				continue
			}
			for i3 := 0; i3 < 12; i3++ {
				if pattern[i3][i-x0] {
					fmt.Fprint(file, "*** ")
				} else {
					fmt.Fprint(file, "    ")
				}

			}
			fmt.Fprintln(file)
		}
	} else if termboxstatus == 1 {
		err := termbox.Init()
		if err != nil {
			return
		}
		defer termbox.Close()
		//w, h := termbox.Size()
		pos := 0
		for i := 0; i < 21; i++ {
			if i == 20 {
				for i6 := 0; i6 < 12; i6++ {
					drawnum(1, 0, data[i6].amount, pos, i)
				}
				continue
			}
			for i5 := 0; i5 < 12; i5++ {
				drawline(1, 0, pattern[i5][i], pos, i)
			}
			pos += 3
		}
		pos = 0
		time.Sleep(10 * time.Second)
	}
}
func drawline(termboxstatus, isfile int, line bool, cordx, cordy int) {
	if termboxstatus == 0 && isfile == 0 {
		if line {
			fmt.Print("***")
		} else {
			fmt.Print("   ")
		}
	} else if termboxstatus == 1 {
		for dr := 0; dr < 3; dr++ {
			if line {
				termbox.SetCell(cordx+dr, cordy, ' ', termbox.ColorGreen, termbox.ColorGreen)
				termbox.Flush()
			} else {
				termbox.SetCell(cordx+dr, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
				termbox.Flush()
			}
		}
		termbox.SetCell(cordx+3, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
		termbox.Flush()
	}
}
func drawnum(termboxstatus, isfile int, num int, cordx, cordy int) {
	if termboxstatus == 0 && isfile == 0 {
		switch num {
		case 0:
			fmt.Print("0  ")
			break
		case 1:
			fmt.Print("1  ")
			break
		case 2:
			fmt.Print("2  ")
			break
		case 3:
			fmt.Print("3  ")
			break
		case 4:
			fmt.Print("4  ")
			break
		case 5:
			fmt.Print("5  ")
			break
		case 6:
			fmt.Print("6  ")
			break
		case 7:
			fmt.Print("7  ")
			break
		case 8:
			fmt.Print("8  ")
			break
		case 9:
			fmt.Print("9  ")
			break
		case 10:
			fmt.Print("10 ")
			break
		case 11:
			fmt.Print("11 ")
			break
		case 12:
			fmt.Print("12 ")
			break
		case 13:
			fmt.Print("13 ")
			break
		case 14:
			fmt.Print("14 ")
			break
		case 15:
			fmt.Print("15 ")
			break
		case 16:
			fmt.Print("16 ")
			break
		case 17:
			fmt.Print("17 ")
			break
		case 18:
			fmt.Print("18 ")
			break
		case 19:
			fmt.Print("19 ")
			break
		case 20:
			fmt.Print("20 ")
			break
		}
	} else if termboxstatus == 1 {
		switch num {
		case 0:
			termbox.SetCell(cordx+0, cordy, '0', termbox.ColorWhite, termbox.ColorBlack)
			termbox.SetCell(cordx+1, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.SetCell(cordx+2, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.SetCell(cordx+3, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.Flush()
			break
		case 1:
			termbox.SetCell(cordx+0, cordy, '1', termbox.ColorWhite, termbox.ColorBlack)
			termbox.SetCell(cordx+1, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.SetCell(cordx+2, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.SetCell(cordx+3, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.Flush()
			break
		case 2:
			termbox.SetCell(cordx+0, cordy, '2', termbox.ColorWhite, termbox.ColorBlack)
			termbox.SetCell(cordx+1, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.SetCell(cordx+2, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.SetCell(cordx+3, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.Flush()
			break
		case 3:
			termbox.SetCell(cordx+0, cordy, '3', termbox.ColorWhite, termbox.ColorBlack)
			termbox.SetCell(cordx+1, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.SetCell(cordx+2, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.SetCell(cordx+3, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.Flush()
			break
		case 4:
			termbox.SetCell(cordx+0, cordy, '4', termbox.ColorWhite, termbox.ColorBlack)
			termbox.SetCell(cordx+1, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.SetCell(cordx+2, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.SetCell(cordx+3, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.Flush()
			break
		case 5:
			termbox.SetCell(cordx+0, cordy, '5', termbox.ColorWhite, termbox.ColorBlack)
			termbox.SetCell(cordx+1, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.SetCell(cordx+2, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.SetCell(cordx+3, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.Flush()
			break
		case 6:
			termbox.SetCell(cordx+0, cordy, '6', termbox.ColorWhite, termbox.ColorBlack)
			termbox.SetCell(cordx+1, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.SetCell(cordx+2, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.SetCell(cordx+3, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.Flush()
			break
		case 7:
			termbox.SetCell(cordx+0, cordy, '7', termbox.ColorWhite, termbox.ColorBlack)
			termbox.SetCell(cordx+1, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.SetCell(cordx+2, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.SetCell(cordx+3, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.Flush()
			break
		case 8:
			termbox.SetCell(cordx+0, cordy, '8', termbox.ColorWhite, termbox.ColorBlack)
			termbox.SetCell(cordx+1, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.SetCell(cordx+2, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.SetCell(cordx+3, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.Flush()
			break
		case 9:
			termbox.SetCell(cordx+0, cordy, '9', termbox.ColorWhite, termbox.ColorBlack)
			termbox.SetCell(cordx+1, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.SetCell(cordx+2, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.SetCell(cordx+3, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.Flush()
			break
		case 10:
			termbox.SetCell(cordx+0, cordy, '1', termbox.ColorWhite, termbox.ColorBlack)
			termbox.SetCell(cordx+1, cordy, '0', termbox.ColorWhite, termbox.ColorBlack)
			termbox.SetCell(cordx+2, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.SetCell(cordx+3, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.Flush()
			break
		case 11:
			termbox.SetCell(cordx+0, cordy, '1', termbox.ColorWhite, termbox.ColorBlack)
			termbox.SetCell(cordx+1, cordy, '1', termbox.ColorWhite, termbox.ColorBlack)
			termbox.SetCell(cordx+2, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.SetCell(cordx+3, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.Flush()
			break
		case 12:
			termbox.SetCell(cordx+0, cordy, '1', termbox.ColorWhite, termbox.ColorBlack)
			termbox.SetCell(cordx+1, cordy, '2', termbox.ColorWhite, termbox.ColorBlack)
			termbox.SetCell(cordx+2, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.SetCell(cordx+3, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.Flush()
			break
		case 13:
			termbox.SetCell(cordx+0, cordy, '1', termbox.ColorWhite, termbox.ColorBlack)
			termbox.SetCell(cordx+1, cordy, '3', termbox.ColorWhite, termbox.ColorBlack)
			termbox.SetCell(cordx+2, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.SetCell(cordx+3, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.Flush()
			break
		case 14:
			termbox.SetCell(cordx+0, cordy, '1', termbox.ColorWhite, termbox.ColorBlack)
			termbox.SetCell(cordx+1, cordy, '4', termbox.ColorWhite, termbox.ColorBlack)
			termbox.SetCell(cordx+2, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.SetCell(cordx+3, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.Flush()
			break
		case 15:
			termbox.SetCell(cordx+0, cordy, '1', termbox.ColorWhite, termbox.ColorBlack)
			termbox.SetCell(cordx+1, cordy, '5', termbox.ColorWhite, termbox.ColorBlack)
			termbox.SetCell(cordx+2, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.SetCell(cordx+3, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.Flush()
			break
		case 16:
			termbox.SetCell(cordx+0, cordy, '1', termbox.ColorWhite, termbox.ColorBlack)
			termbox.SetCell(cordx+1, cordy, '6', termbox.ColorWhite, termbox.ColorBlack)
			termbox.SetCell(cordx+2, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.SetCell(cordx+3, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.Flush()
			break
		case 17:
			termbox.SetCell(cordx+0, cordy, '1', termbox.ColorWhite, termbox.ColorBlack)
			termbox.SetCell(cordx+1, cordy, '7', termbox.ColorWhite, termbox.ColorBlack)
			termbox.SetCell(cordx+2, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.SetCell(cordx+3, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.Flush()
			break
		case 18:
			termbox.SetCell(cordx+0, cordy, '1', termbox.ColorWhite, termbox.ColorBlack)
			termbox.SetCell(cordx+1, cordy, '8', termbox.ColorWhite, termbox.ColorBlack)
			termbox.SetCell(cordx+2, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.SetCell(cordx+3, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.Flush()
			break
		case 19:
			termbox.SetCell(cordx+0, cordy, '1', termbox.ColorWhite, termbox.ColorBlack)
			termbox.SetCell(cordx+1, cordy, '9', termbox.ColorWhite, termbox.ColorBlack)
			termbox.SetCell(cordx+2, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.SetCell(cordx+3, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.Flush()
			break
		case 20:
			termbox.SetCell(cordx+0, cordy, '2', termbox.ColorWhite, termbox.ColorBlack)
			termbox.SetCell(cordx+1, cordy, '0', termbox.ColorWhite, termbox.ColorBlack)
			termbox.SetCell(cordx+2, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.SetCell(cordx+3, cordy, ' ', termbox.ColorBlack, termbox.ColorBlack)
			termbox.Flush()
			break
		}
	}
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
	//drawpattern(0, 0, 0, createpattern("down"))
	//drawpattern(0, 1, 1, createpattern("down"))
}
func task6() {
	//drawpattern(0, 1, 2, createpattern("up"))
	drawpattern(1, 0, 0, createpattern("up"))
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
	if false {
		task5()
	}
	if true {
		task6()
	}
}
