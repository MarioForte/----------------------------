//go:build lab2_5

package main

import "fmt"

type Rectangle struct {
	a, b int
}

func square(rect Rectangle) int {
	return rect.a * rect.b
}

func main() {
	rect1 := Rectangle{13, 15}
	rect2 := Rectangle{14, 27}
	fmt.Printf("Площадь первого прямоугольника: %d\n", square(rect1))
	fmt.Printf("Площадь второго прямоугольника: %d\n", square(rect2))

}
