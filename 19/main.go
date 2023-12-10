package main

/*
Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.
*/

import (
	"fmt"
	"strings"
)

func reverseString(str string) string {
	// преобразовываем строку в слайс рун
	// т.к. строка это []byte, и unicode символы могут быть > 1 байта
	runeSlice := []rune(str)
	var b strings.Builder

	// проходимся по слайсу рун задом наперёд
	for i := len(runeSlice) - 1; i >= 0; i-- {
		b.WriteRune(runeSlice[i])
	}

	return b.String()
}

func main() {
	fmt.Println(reverseString("главрыба"))
}
