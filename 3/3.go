//go:build lab3_3

package main

import (
	"fmt"
	"stringutils"
)

func main() {
	var s string
	fmt.Print("Введите строку для разворота: ")
	fmt.Scan(&s)

	fmt.Println(stringutils.Reverse(s))
}
