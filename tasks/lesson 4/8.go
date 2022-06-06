package main

import (
	"fmt"
	//"strconv"
)

func main() {
	var num int
	//var rez string
	var rez2 int
	fmt.Print("num: ")
	fmt.Scan(&num)
	fmt.Println()
	/*for i := lengthS; i > 0; i-- {
		rez += strconv.Itoa((num % 10))
		num /= 10
	}*/
	for {
		rez2 += (num % 10)
		num /= 10
		if num == 0 {
			break
		}
	}
	fmt.Println(rez2)
}
