package main

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

import (
	"fmt"
)

func main() {
	arr := []int{234234, 2342, 4, 8845, 845, 469, 99, 4556215, 21, 41, 2341, 212, 42131, 624, 6742, 42, 642, 724, 74, 4256, 24634, 5, 25, 4568, 45, 8458, 6548, 58456, 85468}
	QuickSort(arr)
	fmt.Println(arr)
}

func QuickSort(arr []int) {
	low := 0
	high := len(arr) - 1
	quickSort(arr, low, high)
}

func quickSort(arr []int, low int, high int) {

	if low < high {
		partitionIndex := partion(arr, low, high)

		quickSort(arr, low, partitionIndex-1)
		quickSort(arr, partitionIndex+1, high)
	}
}

func partion(arr []int, low int, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++

			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
