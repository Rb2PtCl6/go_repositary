package main

import (
	"fmt"
)

// substr возвращает подстроку строки s, начиная с
// символа #from и заканчивая символом #to
func substr(s string, from, to int) string {
	r := []rune(s)
	return string(r[from : to+1])
}

func main() {
	//                                                   ожидается:
	fmt.Printf("%q\n", substr("abcdefghijklmn", 5, 7))   //  fgh
	fmt.Printf("%q\n", substr("abcdefghijklmn", 10, 17)) //  klmn
	fmt.Printf("%q\n", substr("abcdefghijklmn", 15, 18)) //  пустая строка
	fmt.Printf("%q\n", substr("abcdefghijklmn", 5, 4))   //  пустая строка
	fmt.Printf("%q\n", substr("abcdefghijklmn", 0, 0))   //  a
	fmt.Printf("%q\n", substr("abcdefghijklmn", 0, -1))  //  пустая строка
	fmt.Printf("%q\n", substr("abcdefghijklmn", 5, 3))   //  runtime error
}

/*   Реальный вывод
"fgh"
"klmn\x00\x00\x00\x00"
"\x00\x00\x00\x00"
""
"a"
""
panic: runtime error: slice bounds out of range

goroutine 1 [running]:
main.substr(...)
	H:/samples/11_1.go:11
main.main()
	H:/samples/11_1.go:22 +0x611
exit status 2
*/
