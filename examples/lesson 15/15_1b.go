package main

import (
	"fmt"
	"os"
	"time"

	"github.com/nsf/termbox-go"
)

func putCell(x, y int, c rune, fg, bg termbox.Attribute) {
	termbox.SetCell(x, y, c, fg, bg)
	termbox.Flush()
	time.Sleep(10 * time.Millisecond)
}

func main() {
	if err := termbox.Init(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer termbox.Close()

	width, height := termbox.Size()
	square := width * height
	w, h := width-1, height-2
	i := 0
	x, y := -1, 0

	// (0, 0) - левый верхний угол массива
out:
	for {
		//  идём направо
		for step := 0; step < w; step++ {
			x++
			i++
			putCell(x, y, '─', termbox.ColorRed, termbox.ColorGreen)
			if i == square {
				break out
			}
		}
		x++
		i++
		putCell(x, y, '┐', termbox.ColorRed, termbox.ColorGreen)
		if i == square {
			break
		}
		w--

		//  идём вниз
		for step := 0; step < h; step++ {
			y++
			i++
			putCell(x, y, '│', termbox.ColorRed, termbox.ColorGreen)
			if i == square {
				break out
			}
		}
		y++
		i++
		putCell(x, y, '┘', termbox.ColorRed, termbox.ColorGreen)
		if i == square {
			break
		}
		h--

		//  идём налево
		for step := 0; step < w; step++ {
			x--
			i++
			putCell(x, y, '─', termbox.ColorRed, termbox.ColorGreen)
			if i == square {
				break out
			}
		}
		x--
		i++
		putCell(x, y, '└', termbox.ColorRed, termbox.ColorGreen)
		if i == square {
			break
		}
		w--

		//  идём вверх
		for step := 0; step < h; step++ {
			y--
			i++
			putCell(x, y, '│', termbox.ColorRed, termbox.ColorGreen)
			if i == square {
				break out
			}
		}
		y--
		i++
		putCell(x, y, '┌', termbox.ColorRed, termbox.ColorGreen)
		if i == square {
			break
		}
		h--
	}
	time.Sleep(3 * time.Second)

}
