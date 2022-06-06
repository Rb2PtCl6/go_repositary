package main

import (
	"fmt"
)

func main() {
	var max, min, sum, i, average int
	i = 1
	for ; ; i++ {
		fmt.Print("num: ")
		var num int
		fmt.Scan(&num)
		fmt.Println()
		if i == 1 {
			max = num
			min = num
			sum += num
		} else {
			if num > max {
				max = num
			}
			if num < min {
				min = num
			}
			sum += num
		}
		//i считает введёные числа
		//
		//в первый раз
		//введённое записывается в мах
		//введённое записывается в мин
		//введенное добавляется в сумму
		//
		//последующие разы
		//сравнивается мин с числом, если меньше перезапись мин
		//сравнивается с мах, если больше перезапись мах
		//для получения средного sum делить на i
		//
		//в конце
		//вывод мах
		//вывод мин
		//вывод сум
		//вывод среднего арифметического
		fmt.Println("Continue? y/n")
		var cn string
		fmt.Scan(&cn)
		if cn == "n" {
			average = sum / i
			fmt.Printf("max: %v\n", max)         //max
			fmt.Printf("min: %v\n", min)         //min
			fmt.Printf("sum: %v\n", sum)         //sum
			fmt.Printf("average: %v\n", average) //average
			break
		}
	}
}
