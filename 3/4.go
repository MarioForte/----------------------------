//go:build lab3_4

package main

import (
	"fmt"
)

func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	for _, a := range arr {
		fmt.Println(a)
	}
}
