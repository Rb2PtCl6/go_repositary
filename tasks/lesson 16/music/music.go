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
	max := max_num()
	base = make([]tMusic, max)
	i := 0
	skip := 0
	var ref string
	for reader.Scan() {
		str := []rune(reader.Text())
		var strT0 string
		plus := 0
		for i0 := 0; i0 < 5; i0++ {
			fmt.Println(i0)
			//for {
			for ind, simb := range str {
				if skip > 0 {
					skip--
					continue
				}
				var case0 int
				fmt.Println("cic")
				if str[plus+ind] == '[' || case0 == 0 {
					if ind == 0 {
						case0 = 0
						continue
					}
					strT0 += string(simb)
					//fmt.Println(strT0)
					if str[plus+ind+1] == ']' {
						yu := 0
						fmt.Println("tit")
						//strT1 := []rune(str)
						//strT0 = strT0[:len(strT0)-]
					out:
						for {
							fmt.Println("rir", yu)
							yu++
							var rez string
							strT2 := make([]rune, 1)
							//fmt.Println(string(strT1))
							switch string(strT0) {
							case "artist":
								for i = plus + ind + 2; ; i++ {
									//fmt.Println("lolo")
									if str[i+1] == '[' || str[i+1] == '&' || str[i+1] == '#' {
										rez = string(strT2)
										plus += (i + 2) - plus - ind
										skip = 2
										strT0 = ""
										break
									}
									strT2 = append(strT2, str[i])
									ref = string(strT2)
									fmt.Println(ref)
								}
								base[i].artist = rez
								fmt.Println(base[i].artist)
								case0 = 1
								break out
							case "genre":
								for i = plus + ind + 1; ; i++ {
									if str[i+1] == '[' || str[i+1] == '&' || str[i+1] == '#' {
										rez = string(strT2)
										plus += (i + 2) - plus - ind
										skip = 2
										strT0 = ""
										break
									}
									strT2 = append(strT2, str[i])
									ref = string(strT2)
									fmt.Println(ref)
								}
								base[i].genre = rez
								case0 = 1
								break out
							case "title":
								for i = plus + ind + 1; ; i++ {
									if str[i+1] == '[' || str[i+1] == '&' || str[i+1] == '#' {
										rez = string(strT2)
										plus += (i + 2) - plus - ind
										skip = 2
										strT0 = ""
										break
									}
									strT2 = append(strT2, str[i])
									ref = string(strT2)
									fmt.Println(ref)
								}
								base[i].album_title = rez
								case0 = 1
								break out
							}
						}
					}
				}
				if str[plus+ind] == '#' {
					base[i].year = int(str[plus+ind+2] + str[plus+ind+3] + str[plus+ind+4] + str[plus+ind+5])
					plus += 7
				}
				if str[plus+ind] == '&' {
					if str[plus+ind+4] == ' ' {
						base[i].song_amount = int(str[plus+ind+3])
						plus += 4
					} else {
						base[i].song_amount = int(str[plus+ind+3] + str[plus+ind+4])
						plus += 4
					}
				}
			}
			//continue
			//}
		}
		if i == max-2 {
			break
		}
		i++
		fmt.Println("cicle")
	}
	return
}
func main() {
	//var some float64
	//fmt.Print("text: ")
	//fmt.Scan(&some)
	base := read_file()
	fmt.Println(base[1], "com")
}
