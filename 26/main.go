package main

/*
Разработать программу, которая проверяет, что все символы в строке уникальные
(true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false
*/

import (
	"fmt"
	"strings"
)

func isSet(str string) bool {
	str = strings.ToLower(str)

	runeSlice := []rune(str)
	set := make(map[string]interface{})

	for _, char := range runeSlice {
		if _, ok := set[string(char)]; ok {
			continue
		}
		set[string(char)] = struct{}{}
	}

	return len(runeSlice) == len(set)
}

func main() {
	fmt.Println(isSet("abcd"))
	fmt.Println(isSet("abCdefAaf"))
	fmt.Println(isSet("aabcd"))
}
