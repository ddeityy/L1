package main

import "fmt"

/*
Поменять местами два числа без создания временной переменной.
*/

func main() {
	a, b := 10, 20
	fmt.Printf("a = %d, b = %d\n", a, b)

	a, b = b, a
	fmt.Printf("a = %d, b = %d\n", a, b)
}
