package main

import (
	"bufio"
	"fmt"
	"os"
)

func give7() (string, int) { return "tmp/t7.txt", 8 }
func write7() error {
	name, lenA := give7()
	fin7, err := os.Create(name)
	num_a := make([]int, lenA)
	if err != nil {
		return err
	}
	defer fin7.Close()
	for i := 0; i < len(num_a); i++ {
		fmt.Printf("Number %v: ", i)
		fmt.Scan(&num_a[i])
		fmt.Println()
	}
	fin7_buff := bufio.NewWriter(fin7)
	defer fin7_buff.Flush()
	for i := 0; i < len(num_a); i++ {
		fmt.Fprintf(fin7_buff, "%v ", num_a[i])
	}
	fmt.Fprintln(fin7_buff)
	return nil
}

func read7() ([]int, error) {
	name, lenA := give7()
	fin7, err := os.Create(name)
	num_a := make([]int, lenA)
	if err != nil {
		return []int{}, err
	}
	defer fin7.Close()
	fin7_buff := bufio.NewWriter(fin7)
	defer fin7_buff.Flush()
	//
	return num_a, nil
}

/*func count7() error{
	name, lenA:=give7()
	fin7, err := os.Create(name)
	//num_a := make([]int, lenA)
	if err != nil {
		return err
	}
	defer fin7.Close()
	return nil
}*/
func main() {
	err := write7()
	fmt.Println(err)
	slice1, err1 := read7()
	fmt.Println(slice1)
	fmt.Println(err1)
	/*num_b := make([]int, 8)
	for i := 0; i < len(num_b); i++ {
		fmt.Fscan(fin7, num_b[i])
	}
	for i := 0; i < len(num_b); i++ {
		fmt.Println(num_b)
	}*/
}
