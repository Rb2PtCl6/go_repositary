package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func max_len() (max int) {
	file1, err := os.Open("2.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file1.Close()
	row_s1 := bufio.NewScanner(file1)
	for row_s1.Scan() {
		max++
	}
	return
}
func scanning1() ([]map[string]string, int) {
	file, err := os.Open("2.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	row_s := bufio.NewScanner(file)
	i := 0
	base := make([]map[string]string, max_len())
	for row_s.Scan() {
		base[i] = make(map[string]string)
		str := row_s.Text()
		str1 := strings.Split(str, " ")
		base[i]["who"] = str1[0]
		base[i]["what"] = str1[1]
		i++
	}
	return base, max_len()
}
func count_A_B() (cust, order int) {
	base, max := scanning1()
	cust = max
	order = max
	for i := 0; i < max; i++ {
		for i1 := 0; i1 < i; i1++ {
			if i > 0 {
				if base[i1]["who"] == base[i]["who"] {
					cust--
				}
			}
		}
	}
	for i := 0; i < max; i++ {
		for i1 := 0; i1 < i; i1++ {
			if i > 0 {
				if base[i1]["what"] == base[i]["what"] {
					order--
				}
			}
		}
	}
	return
}
func count_D() (how_m_who map[string]int) {
	base, max := scanning1()
	how_m_who = make(map[string]int)
	for i := 0; i < max; i++ {
		if i == 0 {
			how_m_who[base[i]["who"]]++
			continue
		}
		for i1 := 0; i1 < len(how_m_who); i1++ {
			if how_m_who[base[i]["who"]] > 0 {
				how_m_who[base[i]["who"]]++
				break
			}
			if i1 == len(how_m_who)-1 {
				how_m_who[base[i]["who"]]++
			}
		}
	}
	return
}
func count_G() {
	base, max := scanning1()
	file, err := os.Create("2-rez.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	row_s1 := bufio.NewWriter(file)
	defer row_s1.Flush()
	for i := 0; i < max; i++ {
		fmt.Fprintf(row_s1, "%s %s", string(base[i]["who"]), string(base[i]["what"]))
		if i < max-1 {
			fmt.Fprint(row_s1, "\n")
		}
	}
}
func main() {
	cust, order := count_A_B()
	fmt.Printf("cust: %d\norder: %d\n", cust, order)
	how_m_who := count_D()
	fmt.Println(how_m_who)
	count_G()
}
