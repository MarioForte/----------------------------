//go:build lab1_4

package main

import "fmt"

func main() {
	var a, b float64

	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)

	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)

	fmt.Printf("\nРезультаты операций:\n")
	fmt.Printf("%g + %g = %g\n", a, b, a+b)
	fmt.Printf("%g - %g = %g\n", a, b, a-b)
	fmt.Printf("%g * %g = %g\n", a, b, a*b)

	if b != 0 {
		fmt.Printf("%g / %g = %g\n", a, b, a/b)
		fmt.Printf("%d %% %d = %d\n", int(a), int(b), int(a)%int(b))
	} else {
		fmt.Println("Деление на ноль невозможно!")
	}
}
