package main

/*
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
Приведите корректный пример реализации.

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func main() {
	someFunc()
}
*/

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func createHugeString(size int) string {
	// используем буфер для эффективной конкатенации строк
	var b strings.Builder

	for i := 0; i < size; i++ {
		fmt.Fprint(&b, "的")
	}

	return b.String()
}

var justString string

func someFunc() {
	v := createHugeString(1 << 10) // 1024

	// руна может занимать больше одного байта
	fmt.Println(utf8.RuneLen('的')) // 3
	fmt.Println(utf8.RuneLen('Ж')) // 2
	fmt.Println(utf8.RuneLen('A')) // 1

	// в данном случае мы срезаем по количеству байт, а не по количеству рун
	justString = v[:3]
	fmt.Println(justString) // 的
	justString = v[:2]
	fmt.Println(justString) // пытаемся прочитать 2/3 байт и получаем �

	// преобразовываем строку в слайс рун
	r := []rune(v)

	// в даннам случае мы срезаем по количеству рун
	justString = string(r[:2])
	fmt.Println(justString) // 的的
}
func main() {
	someFunc()
}
