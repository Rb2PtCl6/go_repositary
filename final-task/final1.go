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
func drawpattern(istermbox, isfile, pos int, pattern [12][20]bool) {
	data := reader()
	var x0 int
	if pos == 1 {
		x0 = 1
	}
	if istermbox == 0 && isfile == 0 {
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
	} else if istermbox == 0 && isfile == 1 {
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
					fmt.Fprintf(file, "%-4d", data[i2].amount)
				}
				fmt.Fprintln(file)
				continue
			} else if pos == 2 && i == 20 {
				for i2 := 0; i2 < 12; i2++ {
					fmt.Fprintf(file, "%-4d", data[i2].amount)
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
	} else if istermbox == 1 {
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
					pos += 3
				}
				continue
			}
			for i5 := 0; i5 < 12; i5++ {
				drawline(1, 0, pattern[i5][i], pos, i)
				pos += 3
			}
			pos = 0
		}

		time.Sleep(30 * time.Second)
	}
}
func drawline(istermbox, isfile int, isdraw bool, cordx, cordy int) {
	if istermbox == 0 && isfile == 0 {
		if isdraw {
			fmt.Print("***")
		} else {
			fmt.Print("   ")
		}
	} else if istermbox == 1 {
		for dr := 0; dr < 3; dr++ {
			if isdraw {
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
func drawnum(istermbox, isfile int, num int, cordx, cordy int) {
	if istermbox == 0 && isfile == 0 {
		fmt.Printf("%-3d", num)
	} else if istermbox == 1 {
		termbox.SetCursor(cordx, cordy)
		fmt.Printf("%-4d", num)
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
	drawpattern(0, 1, 1, createpattern("down"))
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
