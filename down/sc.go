package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	file, err := os.Open("d.txt")
	if err != nil {
		fmt.Println("Something went wrong!")
		os.Exit(1)
	}
	i := 1
	name_a := "c3/c3-"
	ext := ".mp4"
	app := "yt-dlp"
	arg0 := "-o"
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		arg1 := name_a + strconv.Itoa(i) + ext
		arg2 := scanner.Text()
		cmd := exec.Command(app, arg0, arg1, arg2)
		stdout, err := cmd.Output()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(string(stdout))
		i++
	}

}
