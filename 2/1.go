//go:build lab2_1

package main

import "fmt"

func main() {
	var a int

	fmt.Print("Введите число: ")
	fmt.Scan(&a)

	if a%2 != 0 {
		fmt.Println("Число нечетное")
	} else {
		fmt.Println("Число четное")
	}
}
