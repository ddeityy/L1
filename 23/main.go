package main

import "fmt"

/*
Удалить i-ый элемент из слайса.
*/

func deleteElemAppend[T any](arr []T, i int) []T {
	return append(arr[:i], arr[i+1:]...)
}

func deleteElemCopy[T any](arr []T, i int) []T {
	return arr[:i+copy(arr[i:], arr[i+1:])]
}

func deleteElemNoOrder[T any](arr []T, i int) []T {
	arr[i] = arr[len(arr)-1]
	arr = arr[:len(arr)-1]
	return arr[:len(arr)-1]
}

func main() {
	index := 2

	// удаление элемента через append
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Deleting element with index", index, "using append")
	fmt.Printf("Old slice: %v\n", a)
	a = deleteElemAppend(a, index)
	fmt.Printf("New slice: %v\n\n", a)

	// удаление элемента через copy
	a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Deleting element with index", index, "using copy")
	fmt.Printf("Old slice: %v\n", a)
	a = deleteElemCopy(a, index)
	fmt.Printf("New slice: %v\n\n", a)

	// удаление элемента без сохранения порядка
	a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Deleting element with index", index, "without preserving order")
	fmt.Printf("Old slice: %v\n", a)
	a = deleteElemNoOrder(a, index)
	fmt.Printf("New slice: %v\n\n", a)
}
