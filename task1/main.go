package main

import (
	"fmt"
)

func main() {
	var a, b, s float64
	var err error

	fmt.Println("Программа вычисления площади прямоугольника по двум заданным сторонам.")
	fmt.Print("Введите размер первой стороны: ")
	_, err = fmt.Scanln(&a)

	if err != nil || a < 0 {
		fmt.Println("Введено недопустомое числовое значение. Программа завершена")
		return
	}

	fmt.Print("Введите размер второй стороны: ")
	_, err = fmt.Scanln(&b)

	if err != nil || b < 0 {
		fmt.Println("Введено недопустомое числовое значение. Программа завершена")
		return
	}

	s = a * b
	fmt.Printf("Площадь прямоугольника равна: %f.\n Программа завершена.\n", s)
}
