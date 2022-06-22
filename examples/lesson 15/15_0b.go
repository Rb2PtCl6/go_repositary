package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type tData2 [100][100]int

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
			if c, err := strconv.Atoi(snum); err == nil {
				a[height][width] = c
				width++
			}
		}
		height++
	}

	// Вывод считанного массива на экран
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			fmt.Printf("%4d", a[i][j])
		}
		fmt.Println()
	}
}
