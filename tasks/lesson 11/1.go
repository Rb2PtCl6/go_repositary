package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func scan() string {
	read := bufio.NewReader(os.Stdin)
	str, err := read.ReadString('\n')
	if err != nil && err != io.EOF {
		return "Something went wrong!"
	}
	rez1 := []byte(str)
	str = string(rez1[:len(rez1)-2])
	return str
}
func count() string {
	str := scan()
	if str == "Something went wrong!" {
		return "Something went wrong!"
	}
	r1 := []rune(str)
	nul_count := 0
	for i := 0; i < len(r1); i++ {
		if i == 0 && string(r1[i]) == " " && string(r1[i+1]) != " " {
			nul_count++
		} else if i == len(r1)-1 && string(r1[i-1]) != " " {
			nul_count++
		} else if string(r1[i]) == " " && string(r1[i+1]) != " " {
			nul_count++
		}
		if i == len(r1)-1 && string(r1[i]) == " " {
			return strconv.Itoa((nul_count - 1))
		}
	}
	nul_count++
	return strconv.Itoa(nul_count)
}
func main() {
	fmt.Println(count())
}
