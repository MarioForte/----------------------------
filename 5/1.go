//go:build lab5_1

package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {

	p := Person{"Алексей", 22}
	personInfo(p)
	birthday(p)

}

func personInfo(p Person) {
	fmt.Printf("Имя: %s, возраст: %d\n", p.name, p.age)
}

func birthday(p Person) {
	fmt.Printf("Через год: %d\n", p.age+1)
}
