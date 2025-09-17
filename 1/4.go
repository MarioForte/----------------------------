//go:build lab1_4

package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)

	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)

	fmt.Printf("\nРезультаты операций:\n")
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	fmt.Printf("%d - %d = %d\n", a, b, a-b)
	fmt.Printf("%d * %d = %d\n", a, b, a*b)

	if b != 0 {
		fmt.Printf("%d / %d = %g\n", a, b, float64(a)/float64(b))
		fmt.Printf("%d %% %d = %d\n", a, b, a%b)
	} else {
		fmt.Println("Деление на ноль невозможно!")
	}
}
