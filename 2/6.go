//go:build lab2_6

package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)

	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)

	result := (a + b) / 2

	fmt.Printf("\nСреднее значение: %d\n", result)
}
