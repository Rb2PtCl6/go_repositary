package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, yS1, xS1, att1 float64 //first point
	var x2, y2, yS2, xS2, att2 float64 //second point
	var x3, y3, yS3, xS3, att3 float64 //third point
	//first line
	fmt.Print("x1: ")
	fmt.Scan(&x1)
	fmt.Println()
	fmt.Print("y1: ")
	fmt.Scan(&y1)
	fmt.Println()
	//second line
	fmt.Print("x2: ")
	fmt.Scan(&x2)
	fmt.Println()
	fmt.Print("y2: ")
	fmt.Scan(&y2)
	fmt.Println()
	//third line
	fmt.Print("x3: ")
	fmt.Scan(&x3)
	fmt.Println()
	fmt.Print("y3: ")
	fmt.Scan(&y3)
	fmt.Println()
	//count xS, yS
	xS1 = x2 - x1
	yS1 = y2 - y1
	xS2 = x3 - x2
	yS2 = y3 - y2
	xS3 = x1 - x3
	yS3 = y1 - x3
	//count att
	att1 = math.Sqrt((yS1 * yS1) + (xS1 * xS1))
	att2 = math.Sqrt((yS2 * yS2) + (xS2 * xS2))
	att3 = math.Sqrt((yS3 * yS3) + (xS3 * xS3))
	//att
	att1 = math.Abs(att1)
	att2 = math.Abs(att2)
	att3 = math.Abs(att3)
	//area, per
	var p, area, perim float64
	perim = att1 + att2 + att3
	p = perim / 2
	area = math.Sqrt(p * (p - att1) * (p - att2) * (p - att3))
	fmt.Print("perimetr: ", perim)
	fmt.Println("area: ", area)
}
