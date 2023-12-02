package main

import (
	"fmt"
	"sync"
)

/*
	Написать программу, которая конкурентно рассчитает значение квадратов
	чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

func main() {
	array := [5]int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}

	for _, num := range array {
		wg.Add(1)
		go func(num int) {
			fmt.Println(num * num)
			wg.Done()
		}(num)
	}

	wg.Wait()
}
