//go:build lab3_6

package main

import (
	"fmt"
)

func main() {
	arr := []string{"Маша", "Паша", "Саша", "Даша", "Виталик"}
	var s string
	for _, v := range arr {
		if len(v) > len(s) {
			s = v
		}
	}

	fmt.Println(s)

}
