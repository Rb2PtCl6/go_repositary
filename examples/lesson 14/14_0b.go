package main

import (
	"fmt"
	"os"
	"time"

	"github.com/nsf/termbox-go"
)

func printString(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func printStringCentered(row int, fg, bg termbox.Attribute, msg string) {
	r := []rune(msg)
	length := len(r)
	h, _ := termbox.Size()
	if length > h {
		msg = string(r[:h])
		length = h
	}
	printString((h-length)/2, row, fg, bg, msg)
}

func main() {
	err := termbox.Init()
	if err != nil {
		os.Exit(1) // Ошибка инициализации termbox
	}
	defer termbox.Close()
	w, h := termbox.Size()

	// Очищаем окно, закрашиваем фон зелёным цветом.
	termbox.Clear(0, termbox.ColorGreen)
	termbox.Flush()

	// Печатаем красным по белому.
	printString(w/2, h/2+1,
		termbox.ColorRed, termbox.ColorWhite,
		"It's a beautiful life")
	termbox.Flush()

	// Печатаем ярко-красным по белому.
	printString(w/2, h/2,
		termbox.ColorRed+termbox.AttrBold, termbox.ColorWhite,
		"It's a beautiful life")
	termbox.Flush()

	time.Sleep(time.Second)

	// Печатаем голубым по синему.
	printStringCentered(2,
		termbox.ColorBlue+termbox.AttrReverse, termbox.ColorCyan,
		"What A Wonderful World")
	termbox.Flush()

	time.Sleep(time.Second)

	termbox.SetCursor(0, h-1)
	termbox.Flush()
	// Печатаем белым по чёрному (в windows).
	fmt.Print("What a wonderful day")
	termbox.HideCursor()
	termbox.Flush()
	time.Sleep(3 * time.Second)
}
