//go:build lab1_6

package main

import "fmt"

func main() {
	var a, b, c float64

	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)

	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)

	fmt.Print("Введите третье число: ")
	fmt.Scan(&c)

	result := (a + b + c) / 3

	fmt.Printf("\nСреднее значение: %g", result)
}
