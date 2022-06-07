package main

import (
	"fmt"
)

func detit(num int) (rez string) {
	var i int = 2
	for {
		if num%i == 0 && num != i {
			rez = "число составное"
			break
		} else if num == i {
			rez = "Число простое"
		}
		i++
	}
	return
}
func main() {
	var num int
	fmt.Print("num: ")
	fmt.Scan(&num)
	fmt.Println()
	fmt.Println(detit(num))
	/*
		взять код из 8 задания для проверки введённый номер делить без остатка на и( его с каждым циклом увеличивать на и++) и если
		i=2
		if num%i==0 && num%num!=0{println("число составное")}
		if i==num{break (для for)}
		i++
	*/
}
