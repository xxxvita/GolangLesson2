package main

import (
	"fmt"
	"math"
)

func main() {
	var d, r, s float64
	var err error

	fmt.Println("Программа вычисления диаметра и длины окружности по заданной площади.")
	fmt.Print("Введите площадь окружности: ")
	_, err = fmt.Scanln(&s)

	if err != nil || s < 0 {
		fmt.Println("Введено недопустомое числовое значение. Программа завершена")
		return
	}

	r = math.Sqrt(s / math.Pi)
	d = 2 * r

	fmt.Printf("Диаметр окружности равен: %f\n: ", d)
	fmt.Printf("Длина окружности равен: %f\n: ", d*math.Pi)
}
