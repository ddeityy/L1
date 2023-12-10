package main

/*
Разработать программу, которая в рантайме способна определить тип переменной:
int, string, bool, channel из переменной типа interface{}.
*/

import (
	"fmt"
	"reflect"
)

func main() {
	arr := []interface{}{"hi", 42, func() {}, struct{}{}, true, 45.6, make(chan int)}

	for _, v := range arr {
		v := reflect.ValueOf(v)
		fmt.Printf("'%v' is type '%s'\n", v, v.Kind().String())
	}
}
