package main

import (
	"fmt"
	"sort"
)

/*
Реализовать бинарный поиск встроенными методами языка.
*/

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	sort.Ints(arr)
	targetIndex, err := binarySearch(arr, 14)
	if err != nil {
		panic(err)
	}
	fmt.Println(targetIndex)

}

func binarySearch(arr []int, target int) (int, error) {
	low := 0
	high := len(arr) - 1

	for low <= high {
		//находим середину слайса
		mid := (low + high) / 2

		if arr[mid] < target {
			// сдвигаем левую границу вправо
			low = mid + 1
		} else if arr[mid] > target {
			// сдвигаем правую границу влево
			high = mid - 1
		} else {
			return mid, nil
		}
	}

	return 0, fmt.Errorf("%v not found", target)
}
