package main

import "fmt"

func main() {
	var year int
	fmt.Print("Введите номер года: ")
	fmt.Scan(&year)

	fmt.Print(year, "-й год - год ")
	switch tmp := year % 10; {
	case tmp <= 1:
		fmt.Print("бело")
	case tmp <= 3:
		fmt.Print("черно")
	case tmp <= 5:
		fmt.Print("зелено")
	case tmp <= 7:
		fmt.Print("красно")
	case tmp <= 9:
		fmt.Print("желто")
	}

	switch year % 12 {
	case 1, 6, 7, 8:
		fmt.Print("го ")
	default:
		fmt.Print("й ")
	}

	switch year % 12 {
	case 0:
		fmt.Println("обезьяны")
	case 1:
		fmt.Println("петуха")
	case 2:
		fmt.Println("собаки")
	case 3:
		fmt.Println("свиньи")
	case 4:
		fmt.Println("крысы")
	case 5:
		fmt.Println("коровы")
	case 6:
		fmt.Println("тигра")
	case 7:
		fmt.Println("кота")
	case 8:
		fmt.Println("дракона")
	case 9:
		fmt.Println("змеи")
	case 10:
		fmt.Println("лошади")
	case 11:
		fmt.Println("козы")
	}
}
