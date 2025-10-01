//go:build lab4_5

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var s string

	fmt.Print("Введите числа через пробел: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s = scanner.Text()

	arr := strings.Split(s, " ")

	fmt.Print("Обратный порядок: ")
	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}
