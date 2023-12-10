package main

/*
	Дана последовательность чисел: 2,4,6,8,10.
	Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
*/

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	array := [5]int32{2, 4, 6, 8, 10}
	var squaredSum int32
	wg := sync.WaitGroup{}

	for _, num := range array {
		wg.Add(1)
		go func(num int32) {
			atomic.AddInt32(&squaredSum, num*num)
			wg.Done()
		}(num)
	}

	wg.Wait()
	fmt.Println(squaredSum)
}
