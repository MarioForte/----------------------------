//go:build lab4_1

package main

import (
	"fmt"
)

func main() {
	people := map[string]int{
		"Алексей": 25,
		"Мария":   30,
		"Иван":    22,
		"Ольга":   28,
	}

	printPeople(people)

	people["Петр"] = 35
	people["Анна"] = 27

	printPeople(people)

	fmt.Println(averageAge(people)) // 2 задание

	delete(people, "Иван") // 3 задание

	printPeople(people)
}

func printPeople(people map[string]int) {
	fmt.Println("---------------------")
	for name, age := range people {
		fmt.Printf("%s - %d\n", name, age)
	}
	fmt.Println("---------------------")
}

func averageAge(people map[string]int) float64 { // 2 задание
	if len(people) == 0 {
		return 0.0
	}

	totalAge := 0
	for _, age := range people {
		totalAge += age
	}

	average := float64(totalAge) / float64(len(people))
	return average
}
