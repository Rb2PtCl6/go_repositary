package main

import (
	"fmt"
)

var (
	colors  = [5]string{"бело", "черно", "зелено", "красно", "желто"}
	endings = [12]string{"й", "го", "й", "й", "й", "й", "го", "го", "го", "й", "й", "й"}
	animals = [12]string{
		"обезьяны", "петуха", "собаки", "свиньи", "крысы", "коровы",
		"тигра", "кота", "дракона", "змеи", "лошади", "козы"}
)

func main() {
	var year int
	fmt.Print("Введите номер года: ")
	fmt.Scan(&year)
	fmt.Printf("%d-й год - год %s%s %s\n",
		year, colors[year%10/2], endings[year%12], animals[year%12])
}
