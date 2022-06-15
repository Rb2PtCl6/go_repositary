package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 2000

type TDataArray []int

func CreateFile(filename string) error {
	// Создаём файл filename с числами от 1 до N
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	buff := bufio.NewWriter(file)
	defer buff.Flush()
	for i := 1; i <= N; i++ {
		fmt.Fprintln(buff, i)
	}
	return nil
}

func ReadFile(filename string) (TDataArray, error) {
	var a TDataArray
	// Считываем числа из файла filename
	file, err := os.Open(filename)
	if err != nil {
		return []int{}, err
	}
	defer file.Close()
	buff := bufio.NewReader(file)
	var c int
	for {
		if _, err = fmt.Fscanf(buff, "%d\n", &c); err != nil {
			break
		}
		a = append(a, c)
	}
	return a, nil
}

func WriteReverse(filename string, a TDataArray) error {
	// Выводим слайс a TDataArray в обратном порядке в файл filename
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	buff := bufio.NewWriter(file)
	defer buff.Flush()
	for i := len(a) - 1; i >= 0; i-- {
		fmt.Fprintln(buff, a[i])
	}
	return nil
}

func main() {
	// Создаём файл numbers0.dat
	CreateFile("numbers0.dat")

	// Считываем числа из файла numbers0.dat
	data, _ := ReadFile("numbers0.dat")

	// Выводим их в обратном порядке в файл numbers0.res
	WriteReverse("numbers0.res", data)
}
