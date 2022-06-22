package main

import (
	"bufio"
	"os"

	"github.com/nsf/termbox-go"
)

const (
	alive = 'x'
	dead  = '.'
)

func main() {
	if err := termbox.Init(); err != nil {
		os.Exit(1) // Ошибка инициализации termbox
	}
	defer termbox.Close()

	width, height := termbox.Size()
	width--

	// Считываем начальную позицию из файла life.dat
	field, err := fieldFromFile("life.dat", width, height)
	if err != nil {
		// Ошибка загрузки начального состояния
		termbox.Close()
		os.Exit(2)
	}

	for {
		drawField(field)
		termbox.Flush()
		ev := termbox.PollEvent()
		if ev.Type == termbox.EventKey {
			if ev.Key == termbox.KeyEsc || ev.Key == termbox.KeyCtrlC {
				break
			}
		}
		nextGeneration(field)
	}
}

func fieldSize(field [][]rune) (width, height int) {
	// Поле в памяти мы делаем больше реального, добавляя по одной строке
	// и колонке с каждой стороны для упрощения функции countNeighbours().
	// Реальное поле устроено так:
	// - номер колонки в диапазоне [ 1; width  ],
	// - номер ряда    в диапазоне [ 1; height ].
	return len(field[0]) - 2, len(field) - 2
}

func fieldNew(width, height int) (field [][]rune) {
	height, width = height+2, width+2
	field = make([][]rune, height)
	for row := 0; row < height; row++ {
		field[row] = make([]rune, width)
	}
	return
}

func fieldClear(field [][]rune) {
	width, height := fieldSize(field)
	for row := 1; row <= height; row++ {
		for col := 1; col <= width; col++ {
			field[row][col] = dead
		}
	}
}

func fieldFromFile(fname string, width, height int) (
	field [][]rune, err error,
) {
	field = fieldNew(width, height)
	fieldClear(field)

	f, err := os.Open(fname)
	if err != nil {
		return
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for row := 1; row <= height && s.Scan(); row++ {
		col := 1
		for _, r := range s.Text() {
			if r == alive {
				field[row][col] = alive
			} else {
				field[row][col] = dead
			}
			col++
			if col > width {
				break
			}
		}
	}

	err = s.Err()
	if err != nil {
		field = nil
	}
	return
}

func drawField(field [][]rune) {
	width, height := fieldSize(field)
	fg, bg := termbox.ColorDefault, termbox.ColorDefault
	for row := 1; row <= height; row++ {
		for col := 1; col <= width; col++ {
			termbox.SetCell(col-1, row-1, field[row][col], fg, bg)
		}
	}
}

func countNeighbours(field [][]rune, row, col int) (result int) {
	if field[row-1][col-1] == alive {
		result++
	}
	if field[row-1][col] == alive {
		result++
	}
	if field[row-1][col+1] == alive {
		result++
	}
	if field[row][col-1] == alive {
		result++
	}
	if field[row][col+1] == alive {
		result++
	}
	if field[row+1][col-1] == alive {
		result++
	}
	if field[row+1][col] == alive {
		result++
	}
	if field[row+1][col+1] == alive {
		result++
	}
	return
}

func nextGeneration(field [][]rune) {
	width, height := fieldSize(field)

	work := fieldNew(width, height)
	for row := 1; row <= height; row++ {
		for col := 1; col <= width; col++ {
			v := field[row][col]
			work[row][col] = v
			n := countNeighbours(field, row, col)
			if field[row][col] == alive {
				if (n < 2) || (n > 3) {
					work[row][col] = dead
				}
			} else /* field[row][col] == dead */ {
				if n == 3 {
					work[row][col] = alive
				}
			}
		}
	}

	for row := 1; row <= height; row++ {
		for col := 1; col <= width; col++ {
			field[row][col] = work[row][col]
		}
	}
}
