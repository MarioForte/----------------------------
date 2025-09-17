//go:build lab1_2

package main

import "fmt"

func main() {
	var i int = 10
	var f float64 = 3.14
	var s string = "Hello"
	var b bool = true

	fmt.Printf("int: %d, float: %.2f, string: %s, bool: %t\n", i, f, s, b)
}
