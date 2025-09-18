//go:build lab2_4

package main

import "fmt"

func main() {
	var s string

	fmt.Println("Введите строку")
	fmt.Scan(&s)

	fmt.Println(len(s))
}
