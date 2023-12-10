package main

/*
	Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0
*/

import (
	"fmt"
	"strings"
)

func setBit(n int64, pos uint) (int64, error) {
	if pos > 63 {
		return 0, fmt.Errorf("%v is out of bounds in a 64bit integer", pos)
	}
	fmt.Printf("Setting bit #%v in\n%64b\n", pos, n)
	fmt.Println("Creating a bitmask")
	fmt.Printf("%64b\n", 1<<pos)
	fmt.Println("Bitwise OR to set the bit")
	n |= 1 << pos
	fmt.Printf("%64b\n", n)
	return n, nil
}

func clearBit(n int64, pos uint) (int64, error) {
	if pos > 63 {
		return 0, fmt.Errorf("%v is out of bounds in a 64bit integer", pos)
	}
	fmt.Printf("Clearing bit #%v in\n%64b\n", pos, n)
	fmt.Println("Creating an XOR bitmask")
	fmt.Printf("%64b\n", 1<<pos)
	fmt.Println("Bitwise AND NOT to clear the bit")
	var mask int64 = 1 << pos
	n &^= mask
	fmt.Printf("%64b\n", n)
	return n, nil
}

func main() {
	var num int64 = 9223372036854775807
	num, err := clearBit(num, 33)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(strings.Repeat("-", 64))
	_, err = setBit(num, 33)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(strings.Repeat("-", 64))
}
