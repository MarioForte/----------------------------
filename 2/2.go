//go:build lab2_2

package main

import "fmt"

func main() {
	var a int

	fmt.Print("Введите число: ")
	fmt.Scan(&a)

	if a > 0 {
		fmt.Println("Positive")
	} else if a < 0 {
		fmt.Println("Negative")
	} else if a == 0 {
		fmt.Println("Zero")
	}
}
