package main

import (
	"fmt"
)

type tData2 [100][100]int

func main() {
	var (
		a             tData2
		height, width int
	)

	fmt.Print("Высота: ")
	fmt.Scan(&height)
	fmt.Print("Ширина: ")
	fmt.Scan(&width)
	i := 0
	x, y := -1, 0
	w, h := width, height-1
	square := width * height
	// (0, 0) - левый верхний угол массива
out:
	for {
		//  идём направо
		for step := 0; step < w; step++ {
			x++
			i++
			a[y][x] = i
			if i == square {
				break out
			}
		}
		w--
		//  идём вниз
		for step := 0; step < h; step++ {
			y++
			i++
			a[y][x] = i
			if i == square {
				break out
			}
		}
		h--
		//  идём налево
		for step := 0; step < w; step++ {
			x--
			i++
			a[y][x] = i
			if i == square {
				break out
			}
		}
		w--
		//  идём вверх
		for step := 0; step < h; step++ {
			y--
			i++
			a[y][x] = i
			if i == square {
				break out
			}
		}
		h--
	}

	// Вывод массива на экран
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			fmt.Printf("%4d", a[i][j])
		}
		fmt.Println()
	}
}
