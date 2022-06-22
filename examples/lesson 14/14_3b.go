package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/nsf/termbox-go"
)

const (
	alive = 'x'
	dead  = '.'
)


func drawField(field [][]rune) {
	for row := 0; row < len(field); row++ {
		for col := 0; col < len(field[row]); col++ {
			termbox.SetCell(
				col, row, field[row][col],
				termbox.ColorDefault, termbox.ColorDefault,
			)
		}
	}
	termbox.Flush()
}

func saveField(field [][]rune, fName string) {
	f, err := os.Create(fName)
	if err != nil {
		termbox.Close()
		os.Exit(3)
	}
	defer f.Close()

	for row := 0; row < len(field); row++ {
		fmt.Fprintln(f, string(field[row]))
	}
}

func main() {
	if err := termbox.Init(); err != nil {
		// Ошибка инициализации termbox
		os.Exit(1)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc + termbox.InputMouse)

	width, height := termbox.Size()
	width--

	// Заполняем массив field мёртвыми клетками...
	var field [][]rune
	deadLine := strings.Repeat(string(dead), width)
	for row := 0; row < height; row++ {
		field = append(field, []rune(deadLine))
	}
	// ... и рисуем его
	drawField(field)

	var (
		mx, my   int
		fileName string
	)
	for {
		ev := termbox.PollEvent()
		if ev.Type == termbox.EventKey {
			if ev.Key == termbox.KeyEsc || ev.Key == termbox.KeyCtrlC {
				break
			} else if ev.Key == termbox.KeyCtrlS {
				termbox.Close()
				fmt.Print("Enter file name: ")
				fmt.Scan(&fileName)
				saveField(field, fileName)
				if err := termbox.Init(); err != nil {
					os.Exit(1)
				}
				drawField(field)
			} else {
				// . . .
			}
		} else if ev.Type == termbox.EventMouse {
			if ev.Key == termbox.MouseLeft {
				mx, my = ev.MouseX, ev.MouseY
				if field[my][mx] == alive {
					field[my][mx] = dead
				} else { // field[my]mx] == dead
					field[my][mx] = alive
				}
				termbox.SetCell(
					mx, my, field[my][mx],
					termbox.ColorDefault, termbox.ColorDefault,
				)
				termbox.Flush()
			}
		}
	}
}
