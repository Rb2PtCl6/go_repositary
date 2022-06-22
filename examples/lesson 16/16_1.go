package main

import (
	"fmt"
)

type (
	tPoint struct {
		x, y float64
	}
	tCircle struct {
		O tPoint
		R float64
	}
)

func main() {

	var circle tCircle
	fmt.Println("Enter the coordinates of the center of the circle: ")
	fmt.Scan(&circle.O.x, &circle.O.y)
	fmt.Println("Enter the radius of the circle: ")
	fmt.Scan(&circle.R)

	var ABC struct {
		A, B, C tPoint
	}
	fmt.Println("Enter the coordinates of the triangle first vertex: ")
	fmt.Scan(&ABC.A.x, &ABC.A.y)
	fmt.Println("Enter the coordinates of the triangle second vertex: ")
	fmt.Scan(&ABC.B.x, &ABC.B.y)
	fmt.Println("Enter the coordinates of the triangle third vertex: ")
	fmt.Scan(&ABC.C.x, &ABC.C.y)

	if 	pointInCircle(ABC.A, circle) &&
		pointInCircle(ABC.B, circle) &&
		pointInCircle(ABC.C, circle) {
		fmt.Println("The whole triangle is in a circle")
	} else {
		fmt.Println("The triangle is not located inside the circle")
	}
}

func pointInCircle(p tPoint, c tCircle) bool {
	return sqr(p.x-c.O.x)+sqr(p.y-c.O.y) <= sqr(c.R)
}

func sqr(t float64) float64 {
	return t * t
}
