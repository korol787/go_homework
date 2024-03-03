package main

import (
	"fmt"
	"strings"
)

var input string

func main() {
	fmt.Print("Введите строку: ")
	fmt.Scanln(&input)

	characters := strings.Split(input, "")

	fmt.Println("Слайс символов строки:")
	fmt.Println(characters)
}
