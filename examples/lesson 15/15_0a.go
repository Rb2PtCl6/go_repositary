package main

import (
	"bufio"
	"fmt"
	"os"
)

type tData2 [100][100]int

func main() {
	var (
		a             tData2
		height, width int
	)

	// Чтение данных из файла numbers2D.dat
	fin, _ := os.Open("numbers2D.dat")
	defer fin.Close()
	reader := bufio.NewReader(fin)

	fmt.Fscanf(reader, "%d %d\n", &height, &width)
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			fmt.Fscan(reader, &a[row][col])
		}
	}

	// Вывод считанного массива на экран
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			fmt.Printf("%8d", a[i][j])
		}
		fmt.Println()
	}
}
