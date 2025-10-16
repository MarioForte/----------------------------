//go:build lab5_3

package main

import "fmt"

type Circle struct {
	radius float64
}

type Shape interface {
	Area() float64
}

type Rectangle struct {
	a float64
	b float64
}

func main() {
	c := Circle{5}
	fmt.Printf("Площадь круга: %g\n", c.Area())
	r := Rectangle{5, 4}
	fmt.Printf("Площадь прямоуглоьника: %g\n", r.Area())

	shapes := []Shape{c, r}
	for i, v := range shapes {
		fmt.Printf("Площадь фигуры %d: %g\n", i, v.Area())
	}
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * 3.14
}

func (r Rectangle) Area() float64 {
	return r.a * r.b
}
