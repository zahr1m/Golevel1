package main

import (
	"fmt"
	"math"
)

func main() {
	var s float64
	fmt.Print("Введите площадь круга S: ")
	fmt.Scanln(&s)
	d := 2 * math.Sqrt(s/math.Pi)
	l := math.Pi * d
	fmt.Println("Диаметр круга: ",d)
	fmt.Println("Длина окружности: ",l)
}
