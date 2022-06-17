package main

import (
	"fmt"
)

func main() {
	a := make([]map[string]string, 3)
	for i := 0; i < 3; i++ {
		a[i] = make(map[string]string)
	}
	a[0]["b"] = "1000"
	a[0]["c"] = "city1"
	a[1]["b"] = "1001"
	a[1]["c"] = "city2"
	a[2]["b"] = "1002"
	a[2]["c"] = "city3"
	fmt.Println(a)
}
