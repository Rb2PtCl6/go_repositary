package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/nsf/termbox-go"
)

func draw(maxX, maxY int, color0 termbox.Attribute) {
	file, _ := os.Create("log0.txt")
	defer file.Close()
	row_s1 := bufio.NewWriter(file)
	defer row_s1.Flush()
	for i := 0; i < maxY; i++ {
		sub_draw(maxX, i, color0)
		fmt.Fprintln(row_s1, i, 2222)
	}
}
func sub_draw(maxX, y int, color0 termbox.Attribute) {
	var rune0 rune = ' '
	file, _ := os.Create("log.txt")
	defer file.Close()
	row_s1 := bufio.NewWriter(file)
	defer row_s1.Flush()
	for i1 := 0; i1 < maxX; i1++ {
		termbox.SetCell(i1, y, rune0, color0, color0)
		termbox.Flush()
		fmt.Fprintln(row_s1, rune0, i1, y, color0, 1111)
	}
}
func main() {
	var start, startU string
	start = "yes"
	fmt.Print("if you want to start, type 'yes': ")
	fmt.Scan(&startU)
	if start == startU {
		err := termbox.Init()
		if err != nil {
			os.Exit(1)
		}
		defer termbox.Close()
		w, h := termbox.Size()
		draw(w, h, termbox.ColorRed)
		//fmt.Println(err, w, h)
	}
	time.Sleep(4 * time.Second)
}
