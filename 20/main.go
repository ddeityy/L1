package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow»
*/

func reverseWords(str string) string {
	words := strings.Split(str, " ")

	var b strings.Builder

	for i := len(words) - 1; i >= 0; i-- {
		b.WriteString(words[i])
		b.WriteString(" ")
	}

	result := strings.TrimSpace(b.String())

	return result
}

func main() {
	words := "snow dog sun"
	fmt.Println("initial words:", words)
	fmt.Println("reversed words:", reverseWords(words))
}
