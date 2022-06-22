package main

import (
	"bufio"
	"fmt"
	"os"
)

type (
	tArr [10000]int
	tData struct  {
		nums tArr
		len int
	}	
)

func main() {
	// Считываем числа из файла numbers2.dat

	fin, err := os.Open("numbers.dat")
	if err != nil {
		return
	}
	defer fin.Close()

	in := bufio.NewReader(fin)
	var a tData
	for {
		if _, err = fmt.Fscanf(in, "%d\n", &a.nums[a.len]); err != nil {
			break
		}
		a.len++
	}
	
	fmt.Println(arrMax(a))
}

func arrMax (d tData) (max int) {
	max = d.nums[0]
	for i:= 1; i < d.len; i++  {
		if max < d.nums[i]  { max = d.nums[i] }
	}	
	return
}	
