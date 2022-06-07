package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	time.Sleep(500 * time.Millisecond)
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(elapsed)
}
