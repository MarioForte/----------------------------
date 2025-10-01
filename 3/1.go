//go:build lab3_1

package main

import (
	"fmt"
	"mathutils"
)

func main() {
	var number int

	fmt.Print("Введите число для вычисления факториала: ")
	fmt.Scan(&number)

	result := mathutils.Factorial(number)

	if result == -1 {
		fmt.Println("Ошибка: Факториал отрицательного числа не определен!")
	} else {
		fmt.Printf("%d! = %d\n", number, result)
	}
}
