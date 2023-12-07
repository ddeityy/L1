package main

/*
Разработать программу, которая перемножает, делит, складывает,
вычитает две числовых переменных a,b, значение которых > 2^20.
*/

import (
	"fmt"
	"math/big"
)

func main() {
	a, ok := big.NewInt(0).SetString("4854641245452362623423142566574624898924398", 10)
	if !ok {
		panic("Failed to create a")
	}
	b, ok := big.NewInt(0).SetString("1265341523416431431434163614315345555552222", 10)
	if !ok {
		panic("Failed to create b")
	}

	fmt.Printf("1st number: %s\n2nd number: %s\n\n", a.String(), b.String())

	result := big.NewInt(0)

	fmt.Println("Add:", result.Add(a, b))
	fmt.Println("Substract:", result.Sub(a, b))
	fmt.Println("Multiply:", result.Mul(a, b))
	fmt.Println("Divide:", result.Div(a, b))
}
