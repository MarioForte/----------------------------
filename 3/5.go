//go:build lab3_5

package main

import (
	"fmt"
)

func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	arr2 := arr[0:5]

	arr2 = append(arr2, 6, 7)
	i := 2
	arr2 = append(arr2[:i], arr2[i+1:]...)
	for i, v := range arr2 {
		fmt.Printf("%d: %d\n", i, v)
	}
}
