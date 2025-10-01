//go:build lab4_5

package main

import "fmt"

func main() {
	var a, b, c float64

	fmt.Println("Введите три числа для сложения:")
	fmt.Print("Первое число: ")
	fmt.Scan(&a)

	fmt.Print("Второе число: ")
	fmt.Scan(&b)

	fmt.Print("Третье число: ")
	fmt.Scan(&c)

	sum := a + b + c
	fmt.Printf("Сумма: %.2f + %.2f + %.2f = %.2f\n", a, b, c, sum)
}
