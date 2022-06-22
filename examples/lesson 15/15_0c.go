package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type (
	tData2 [100][100]int
	tData1 [100]int
)

func main() {
	var (
		a             tData2
		height, width int
	)

	// Чтение данных из файла numbers2Da.dat
	fin, _ := os.Open("numbers2Da.dat")
	defer fin.Close()

	for scanner := bufio.NewScanner(fin); scanner.Scan(); {
		if scanner.Text() == "" {
			break
		}
		width = 0
		for _, snum := range strings.Fields(scanner.Text()) {
			c, err := strconv.Atoi(snum)
			if err == nil {
				a[height][width] = c
				width++
			}
		}
		height++
	}

	// Вывод считанного массива на экран
	w := maxWidth(a, height, width) + 1
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			fmt.Printf("%*d", w, a[i][j])
		}
		fmt.Println()
	}
}

func maxWidth(a tData2, height, width int) (max int) {
	for col := 0; col < width; col++ {
		for row := 0; row < height; row++ {
			if w := numWidth(a[row][col]); w > max {
				max = w
			}
		}
	}
	return
}

func numWidth(x int) (width int) {
	if x <= 0 {
		width = 1
	}
	for x != 0 {
		x /= 10
		width++
	}
	return
}
