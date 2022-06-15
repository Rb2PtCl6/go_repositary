package main

import (
	"fmt"
)

func main() {
	fmt.Println(" Эта программа переворачивает введенные строки.")
	for i := 1; i == 1; fmt.Scan(&i) {
		var abc string
		fmt.Println(" Введите строку:")
		fmt.Scan(&abc)
		fmt.Println(Reverse1(abc))
		fmt.Println(Reverse2(abc))
		fmt.Println(Reverse3(abc))
		fmt.Print("   Еще раз? /да-1, нет-0/ ")
	}
}

func Reverse1(s string) (res string) {
	for _, x := range s {
		res = string(x) + res
	}
	return
}

func Reverse2(s string) string {
	test := []rune(s)
	for left, right := 0, len(test)-1; left < right; {
		test[left], test[right] = test[right], test[left]
		left++
		right--
	}
	return string(test)
}

func Reverse3(s string) string {
	r := []rune(s)
	res := make([]rune, len(r))
	for i, k := 0, len(r); i < len(r); i++ {
		k--
		res[k] = r[i]
	}
	return string(res)
}
