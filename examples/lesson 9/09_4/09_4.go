package main

import (
	"fmt"
	"os"
)

// Считываем числа из файла numbers.dat
// Сортируем их и записываем результат в файл numbers.res
func main() {

	// Ввод данных
	fin, _ := os.Open("numbers.dat")
	defer fin.Close()
	var a []int
	for {
		var c int
		if _, err := fmt.Fscanf(fin, "%d\n", &c); err != nil {
			break
		}
		a = append(a, c)
	}

	// Сортировка
	simpleSelectionSort(a)

	// Вывод данных
	fout, _ := os.Create("numbers.res")
	defer fout.Close()
	for _, c := range a {
		fmt.Fprintln(fout, c)
	}

}

// Сортировка простым выбором (наибольшего)
func simpleSelectionSort(a []int) {
	for k := len(a); k > 1; k-- {
		// Ищем, где находится наибольший элемент из a[:k]
		j := maxIndex(a[:k])
		//... и меняем его местами с последним элементом a[:k]
		a[j], a[k-1] = a[k-1], a[j]
	}
}

//  Возвращает номер места, где находится наибольший элемент из data
func maxIndex(data []int) (place int) {
	place = 0
	currentMax := data[0]
	for i, v := range data {
		if v > currentMax {
			place = i
			currentMax = v
		}
	}
	return
}
