//go:build lab5_6

package main

import "fmt"

type Book struct {
	Title  string
	Author string
}

type Stringer interface {
	String() string
}

func (b Book) String() string {
	return fmt.Sprintf("Книга: %s - %s", b.Author, b.Title)
}

func main() {
	b := Book{"Чистый код", "Роберт Мартин"}
	fmt.Println(b.String())
}
