package main

import (
	"bufio"
	"fmt"
	"os"
)

type tMusic struct {
	//example: [artist]Billy Crawford [genre]Pop, R&B ##2002 [title]Ride &&&12
	genre       string //from example: [genre]Pop, R&B
	album_title string //from example: [title]Ride
	artist      string //from example: [artist]Billy Crawford
	year        int    //from example: ##2002
	song_amount int    //from example: &&&12
}

const file_name string = "music.dat"

func max_num() int {
	file1, err := os.Open(file_name)
	if err != nil {
		fmt.Println(err)
	}
	defer file1.Close()
	reader := bufio.NewScanner(file1)
	i := 0
	for reader.Scan() {
		if reader.Text() != "" {
			i++
		}
	}
	return i
}
func read_file() (base []tMusic) {
	file, err := os.Open(file_name)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	reader := bufio.NewScanner(file)
	i := 0
	max := max_num()
	base = make([]tMusic, max)
	i := 0
	for reader.Scan() {
		str := []rune(reader.Text())
		var strT0 []rune
		for i0 := 0; i0 < 5; i0++ {
			for {
				for ind, simb := range str {
					if simb == '[' {
						//
					}
					if simb == '#' {
						base[i].year = int(str[2] + str[3] + str[4] + str[5])
					}
					if simb == '&' {
						base[i].song_amount = int(str[3] + str[4])
						//str=str[]
					}
				}
			}
		}
		if i == max-2 {
			break
		}
		i++
	}
	return
}
func main() {
	//var some float64
	//fmt.Print("text: ")
	//fmt.Scan(&some)
	base := read_file()
	fmt.Println(base[10])
}
