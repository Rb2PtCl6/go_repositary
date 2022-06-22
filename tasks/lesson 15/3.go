package main

import (
	"fmt"
)

func main() {
	var mount string
	moun := [...]string{"январь", "февраль", "март", "апрель", "май", "июнь", "июль", "август", "сентябрь", "октябрь", "ноябрь", "декабрь"}
	fmt.Print("mount: ")
	fmt.Scan(&mount)
	fmt.Println()
	mount1:=[]rune(mount)
	for i,_:=range mount1{
		if i//если сомвол большой, сделать его маленьким( юникод),а затем заменить его для для более точного сравнения
	}
	switch mount {
	case moun[0]:
		fmt.Println("31")
	case moun[1]:
		fmt.Println("28")
	case moun[2]:
		fmt.Println("31")
	case moun[3]:
		fmt.Println("30")
	case moun[4]:
		fmt.Println("31")
	case moun[5]:
		fmt.Println("30")
	case moun[6]:
		fmt.Println("31")
	case moun[7]:
		fmt.Println("31")
	case moun[8]:
		fmt.Println("30")
	case moun[9]:
		fmt.Println("31")
	case moun[10]:
		fmt.Println("30")
	case moun[11]:
		fmt.Println("31")
	}
}
