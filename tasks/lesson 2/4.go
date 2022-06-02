package main

import (
	"fmt"
)

func main() {
	var a, b, c, d float64
	//a
	fmt.Print("a: ")
	fmt.Scan(&a)
	fmt.Println()
	//b
	fmt.Print("b: ")
	fmt.Scan(&b)
	fmt.Println()
	//c
	fmt.Print("c: ")
	fmt.Scan(&c)
	fmt.Println()
	//d
	fmt.Print("d: ")
	fmt.Scan(&d)
	fmt.Println()
	//logic
	if a==b && a==c && a==d && b==c && b==d && c==d {
		fmt.Println(a)
	}else if (a==b && b==d) || (b==c && c==d) || (a==c && b==c) || (a==c && c==d) {
		if 
	}else if ((a==b && c==d) || (b==c && a==d) || (a==c && b==d)) {
		if a!=b {
			fmt.Println(a)
			fmt.Println(b)
		}else{
			fmt.Println(a)
			fmt.Println(c)
		}
	} else if{
		2-3
	} else if{
		3-3-1
	} else if{
		3-3-2
	} else if{
		3-3-3
	}else{
		fmt.Println("Unknown error!")
	}
}