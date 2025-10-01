//go:build lab4_4

package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string

	fmt.Print("Введите строку: ")
	fmt.Scan(&s)

	fmt.Println(strings.ToUpper(s))
}
