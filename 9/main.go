package main

import "fmt"

/*
	Разработать конвейер чисел.
	Даны два канала: в первый пишутся числа (x) из массива,
	во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
*/

func createNumsChan(arr []int) <-chan int {
	outCh := make(chan int, len(arr))
	go func() {
		for _, num := range arr {
			outCh <- num
		}
		close(outCh)
	}()
	return outCh
}

func createSquareChan(nums <-chan int) <-chan int {
	outCh := make(chan int)
	go func() {
		for num := range nums {
			outCh <- num * num
		}
		close(outCh)
	}()
	return outCh
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numCh := createNumsChan(array)
	sqNumCh := createSquareChan(numCh)

	for v := range sqNumCh {
		fmt.Println(v)
	}

}
