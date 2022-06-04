// Пример 01
// Предполагается, что часы в задаются в интервале [0; 12) а минуты в [0; 60)

package main

import (
	"fmt"
	"math"
)

func main() {

	var hours, mins int
	fmt.Print("Введите время\nчасы: ")
	fmt.Scan(&hours)

	fmt.Print("минуты: ")
	fmt.Scan(&mins)

	var m_angle, h_angle float64
	m_angle = float64(mins) * 6.0                     // = float64(mins) / 60.0 * 360.0
	h_angle = float64(hours)*15.0 + float64(mins)*.25 // = float64(hours) / 24 * 360 + float(mins) / 60 * 15

	fmt.Println("Разница в углах стрелок", math.Abs(m_angle-h_angle), "градусов.")
}
