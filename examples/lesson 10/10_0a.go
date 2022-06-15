package main

import (
	"fmt"
)

// читает и печатает строку побайтово
func printBytes(s string) {
	fmt.Print("\t")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%5c", s[i])
	}
	fmt.Println()
	fmt.Print("\t")
	for i := 0; i < len(s); i++ {
		fmt.Printf("|%4d", s[i])
	}
	fmt.Println()
}

// читает и печатает строку посимвольно
func printRunes(s string) {
	fmt.Print("\t")
	for _, c := range s {
		fmt.Printf("%5c", c)
	}
	fmt.Println()
	fmt.Print("\t")
	for _, c := range s {
		fmt.Printf("|%4d", c)
	}
	fmt.Println()
}

func main() {
	names := []string{
		"e.g. or i.e.?", // английский
		"télégraphe",    // французский
		"Rīgas torņi",   // латышский
		"телеграф",      // русский
		"տելեգրաֆ",      // армянский
		"ტელეგრაფი",     // грузинский
		"電報",            // китайский
		"kẹtẹkẹtẹ",      // язык Йоруба
		"[ ]​␠\t␈\b( )", // разные пробелы
		"❣∞¼∛⅀𝞨",
		"☼☁☂😀🇦🇶",
	}
	for _, name := range names {
		fmt.Println(name)
		printBytes(name)
		printRunes(name)
		fmt.Println("0...4...8.0...4...8.0...4...8.0...4...8.0...4...8.0...4...8.0...4...8.0...4...8.")
	}
}
