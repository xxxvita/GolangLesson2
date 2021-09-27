package main

import (
	"fmt"
)

func main() {
	var a int16
	// десятки, сотни, единицы
	var d, s, e int16

	var err error

	fmt.Println("Программа разложения трёхзначного числа на сотни, десятки, единицы.")
	fmt.Print("Введите трёхзначное число: ")
	_, err = fmt.Scanln(&a)

	if err != nil || a < 100 || a > 999 {
		fmt.Println("Введено недопустомое числовое значение. Программа завершена")
		return
	}

	e = a % 10
	a = a / 10

	d = a % 10
	a = a / 10

	s = a

	fmt.Printf("Число сотен: %d, число десяток: %d, число единиц: %d\n", s, d, e)
}
