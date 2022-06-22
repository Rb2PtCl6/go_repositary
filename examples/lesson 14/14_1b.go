package main

import (
	"os"

	"github.com/nsf/termbox-go"
)

func frame() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	w, h := termbox.Size()
	termbox.SetCell(0, 0, '┌', termbox.ColorGreen, termbox.ColorDefault)
	termbox.SetCell(w-1, 0, '┐', termbox.ColorGreen, termbox.ColorDefault)
	termbox.SetCell(0, h-1, '└', termbox.ColorGreen, termbox.ColorDefault)
	termbox.SetCell(w-1, h-1, '┘', termbox.ColorGreen, termbox.ColorDefault)
	for col := 1; col < w-1; col++ {
		termbox.SetCell(col, 0, '─', termbox.ColorGreen, termbox.ColorDefault)
		termbox.SetCell(col, h-1, '─', termbox.ColorGreen, termbox.ColorDefault)
	}
	for row := 1; row < h-1; row++ {
		termbox.SetCell(0, row, '│', termbox.ColorGreen, termbox.ColorDefault)
		termbox.SetCell(w-1, row, '│', termbox.ColorGreen, termbox.ColorDefault)
	}
	termbox.HideCursor()
}

func main() {
	err := termbox.Init()
	if err != nil { // Ошибка инициализации termbox
		os.Exit(1)
	}
	defer termbox.Close()

	// Рисует зелёную рамку по периметру окна.
	// При изменении размеров окна рамка перерисовывается
	frame()
	for {
		termbox.Flush()
		ev := termbox.PollEvent()
		if ev.Type == termbox.EventResize {
			frame()
		} else if ev.Type == termbox.EventKey {
			if ev.Key == termbox.KeyEsc || ev.Key == termbox.KeyCtrlC {
				break
			}
		}
	}
}
