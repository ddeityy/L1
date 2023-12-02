package main

import "fmt"

/*
	Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0
*/

func setBit(n int64, pos uint) int64 {
	n |= (1 << pos)
	return n
}

func clearBit(n int64, pos uint) int64 {
	var mask int64 = ^(1 << pos)
	n &= mask
	return n
}

func main() {
	num := int64(6)
	num = setBit(num, 4)
	fmt.Println(num)
	num = clearBit(num, 2)
	fmt.Println(num)
}
