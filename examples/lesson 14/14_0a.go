package main

import (
	"os"
	"time"

	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil { // Ошибка инициализации termbox
		os.Exit(1)
	}
	defer termbox.Close()

	w, h := termbox.Size()
	termbox.HideCursor()

	//  верхний ряд
	termbox.SetCursor(0, 0)
	for col := 0; col < w; col++ {
		termbox.SetCell(col, 0, '>', termbox.ColorWhite, termbox.ColorRed)
		termbox.Flush()
		time.Sleep(20 * time.Millisecond)
	}

	//  правый край
	for row := 1; row < h; row++ {
		termbox.SetCell(w-1, row, 'V', termbox.ColorBlack, termbox.ColorCyan)
		termbox.Flush()
		time.Sleep(20 * time.Millisecond)
	}

	//  нижний ряд
	for col := w - 1; col > 0; col-- {
		termbox.SetCell(col, h-1, '<', termbox.ColorBlue, termbox.ColorGreen)
		termbox.Flush()
		time.Sleep(20 * time.Millisecond)
	}

	//  левый край
	for row := h - 1; row > 0; row-- {
		termbox.SetCell(0, row, '^', termbox.ColorYellow, termbox.ColorMagenta)
		termbox.Flush()
		time.Sleep(20 * time.Millisecond)
	}

	time.Sleep(3 * time.Second)
}
