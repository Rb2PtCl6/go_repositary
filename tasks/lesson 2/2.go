package main

import (
	"fmt"
)

/*func main() {
	var meter, dijm, funt, jard, mil float64
	fmt.Print("meter: ")
	fmt.Scan(&meter)
	fmt.Println()
	//meter=12
	dijm = meter * 100 / 7.5
	funt = dijm / 12
	jard = funt / 3
	mil = funt / 1760
	//dijm
	fmt.Print("dijm: ")
	fmt.Print(dijm)
	fmt.Println()
	//funt
	fmt.Print("funt: ")
	fmt.Print(funt)
	fmt.Println()
	//jard
	fmt.Print("jard: ")
	fmt.Print(jard)
	fmt.Println()
	//mil
	fmt.Print("mil: ")
	fmt.Print(mil)
	fmt.Println()
}*/
func main() {
	var meter, dijm, funt, jard, mil float64
	fmt.Print("meter: ")
	fmt.Scan(&meter)
	fmt.Println()
	//meter=12
	dijm = meter * 100 / 7.5
	funt = dijm / 12
	jard = funt / 3
	mil = funt / 1760
	var ord int
	fmt.Println("Conv to: (type number)")
	fmt.Println("1) dijm")
	fmt.Println("2) funt")
	fmt.Println("3) jard")
	fmt.Println("4) mil")
	fmt.Print("Conv to: ")
	fmt.Scan(&ord)
	if ord == 1 {
		//dijm
		fmt.Print("dijm: ")
		fmt.Print(dijm)
		fmt.Println()
	} else if ord == 2 {
		//funt
		fmt.Print("funt: ")
		fmt.Print(funt)
		fmt.Println()
	} else if ord == 3 {
		//jard
		fmt.Print("jard: ")
		fmt.Print(jard)
		fmt.Println()
	} else if ord == 4 {
		//mil
		fmt.Print("mil: ")
		fmt.Print(mil)
		fmt.Println()
	} else {
		fmt.Println("Unknown error!")
	}
}
