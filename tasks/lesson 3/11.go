package main

import (
	"fmt"
)

func main() {
	var num int
	name1, name2, name3 := "дорллар", "доллара", "долларов"
	var name string
	fmt.Print("num: ")
	fmt.Scan(&num)
	fmt.Println()
	if (num % 10) == 1 {
		name = name1
	} else if (num%10) == 2 || (num%10) == 3 || (num%10) == 4 {
		name = name2
	} else if (num%10) == 5 || (num%10) == 6 || (num%10) == 7 || (num%10) == 8 || (num%10) == 9 {
		name = name3
	}
	fmt.Printf("%v %v", num, name)
}
