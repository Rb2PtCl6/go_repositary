package main

import (
	"os"

	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		os.Exit(1) // Ошибка инициализации termbox
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc + termbox.InputMouse)

	x, y := termbox.Size()
	x, y = x/2, y/2
	termbox.SetCell(
		x, y, '?', termbox.ColorWhite, termbox.ColorRed+termbox.AttrBold,
	)
	ch := '0'
	for {
		termbox.Flush()
		ev := termbox.PollEvent()
		if ev.Type == termbox.EventKey {
			if ev.Key == termbox.KeyArrowUp {
				termbox.SetCell(x, y, '^',
					termbox.ColorWhite, termbox.ColorBlue+termbox.AttrBold)
			} else if ev.Key == termbox.KeyArrowDown {
				termbox.SetCell(x, y, 'v',
					termbox.ColorWhite, termbox.ColorBlue+termbox.AttrBold)
			} else if ev.Key == termbox.KeyArrowLeft {
				termbox.SetCell(x, y, '<',
					termbox.ColorWhite, termbox.ColorBlue+termbox.AttrBold)
			} else if ev.Key == termbox.KeyArrowRight {
				termbox.SetCell(x, y, '>',
					termbox.ColorWhite, termbox.ColorBlue+termbox.AttrBold)
			} else if ev.Key == termbox.KeyEsc || ev.Key == termbox.KeyCtrlC {
				break
			}
		} else if ev.Type == termbox.EventMouse {
			if ev.Key == termbox.MouseLeft {
				termbox.SetCell(x, y, ch,
					termbox.ColorRed, termbox.ColorGreen)
				ch++
			}
			if ev.Key == termbox.MouseRelease {
				termbox.SetCell(x, y, ch,
					termbox.ColorBlack, termbox.ColorWhite)
			}
		}
	}
}
