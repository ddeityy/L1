package main

/*
	Реализовать пересечение двух неупорядоченных множеств.
*/

import "fmt"

func findIntersection(nums1 []int, nums2 []int) []int {
	intersection := make(map[int]int)
	var result []int
	for _, v := range nums1 {
		intersection[v] += 1
	}
	for _, v := range nums2 {
		intersection[v] += 1
	}
	for k, v := range intersection {
		if v > 1 {
			result = append(result, k)
		}
	}
	return result
}

func main() {
	nums1 := []int{5, 1, 36, 58, 14, 52, 77}
	nums2 := []int{32, 58, 78, 52, 4, 3, 93}
	intersection := findIntersection(nums1, nums2)
	fmt.Println(intersection)
}
