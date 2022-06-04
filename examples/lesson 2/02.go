// Пример 02

package main

import (
	"fmt"
)

func main() {
	var a int = 91
	var b int // равносильно var b int = 0

	var c = 91 // равносильно :=
	d := 91    // тип определяется автоматически
	fmt.Println(a, b, c, d)

	// картежное присваивание
	var e1, e2 int = 91, 19 // можно использовать цифры, но не в начале
	var _e, e_ int = 19, 91 // можно использовать '_', пока есть другие символы
	fmt.Println(e1, e2, _e, e_)

	// так тоже можно, тип определяется для каждой переменной отдельно
	var s, f = "abc", 2.512
	s_, f_ := "abc", 2.512
	fmt.Println(s, f, s_, f_)

	var old int = 1
	// значение old перезаписывается, главное чтоб хотя бы 1 новая переменная была
	old, new := 2, 3
	fmt.Println(old, new)

	// технически можно, но настоятельно не рекомендуется
	var __ int
	var кириллица int
	fmt.Println(__, кириллица)

	// НЕЛЬЗЯ
	// var e int := 91
	// var e := 91
	// var 1e int
	// var _ int
}