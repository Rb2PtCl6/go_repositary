package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var varnt int
	varnt = rand.Intn(601)
	fmt.Println("Загадайте число до 600 включительно")
	for {
		fmt.Printf("Предположение комопьютера - %v\n", varnt)
		fmt.Println("Намного больше (a)")       //+50
		fmt.Println("Немного больше(b)")        //+5
		fmt.Println("Чуть больше(c)")           //+2
		fmt.Println("Самую крошеку больше(d)")  //+1
		fmt.Println("Верно (e)")                //correct
		fmt.Println("Самую крошеку меньше (f)") //-1
		fmt.Println("Чуть меньше (g)")          //-2
		fmt.Println("Немного больше(h)")        //-5
		fmt.Println("Намного меньше (i)")       //-50
		fmt.Print("Ваша реакция: ")
		var react string
		fmt.Scan(&react)
		fmt.Println()
		if react == "a" {
			varnt += 50
		} else if react == "b" {
			varnt += 5
		} else if react == "c" {
			varnt += 2
		} else if react == "d" {
			varnt += 1
		} else if react == "e" {
			fmt.Println("Это было интересно. Спасибо!")
			break
		} else if react == "f" {
			varnt -= 1
		} else if react == "g" {
			varnt -= 2
		} else if react == "h" {
			varnt -= 5
		} else if react == "i" {
			varnt -= 50
		}
	}
}
