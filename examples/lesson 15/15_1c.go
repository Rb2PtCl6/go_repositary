package main

import (
	"fmt"
	"os"
	"time"

	"github.com/nsf/termbox-go"
)

const (
	left  = iota  // left = 0
	down  		  // down = 1
	right 		  // right = 2
	up    		  // up = 3
)

func putCell(x, y int, c rune, fg, bg termbox.Attribute) {
	termbox.SetCell(x, y, c, fg, bg)
	termbox.Flush()
	time.Sleep(5 * time.Millisecond)
}

func main() {
	if err := termbox.Init(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer termbox.Close()

	width, height := termbox.Size()
	square := width * height
	wLen, hLen := width-1, height-2

	// начальное состояние
	counter := 0
	x, y := -1, 0
	direction := right
	var (
		dx, dy       int
		sideLen      int
		side, corner rune
	)
	// (0, 0) - левый верхний угол массива
out:
	for {
		switch direction {
		case right:
			dx, dy = 1, 0
			sideLen = wLen
			wLen--
			side, corner, direction = '─', '┐', down
		case left:
			dx, dy = -1, 0
			sideLen = wLen
			wLen--
			side, corner, direction = '─', '└', up
		case down:
			dx, dy = 0, 1
			sideLen = hLen
			hLen--
			side, corner, direction = '│', '┘', left
		case up:
			dx, dy = 0, -1
			sideLen = hLen
			hLen--
			side, corner, direction = '│', '┌', right
		}

		for i := 0; i < sideLen; i++ {
			counter++
			x, y = x+dx, y+dy
			putCell(x, y, side, termbox.ColorRed, termbox.ColorGreen)
			if counter == square {
				break out
			} // этот if можно убрать
		}

		counter++
		x, y = x+dx, y+dy
		putCell(x, y, corner, termbox.ColorRed, termbox.ColorGreen)
		if counter == square {
			break
		}
	}
	time.Sleep(3 * time.Second)
}
