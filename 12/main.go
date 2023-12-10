package main

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree)
создать для нее собственное множество.
*/

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	m := make(map[string]struct{})
	res := make([]string, 0)

	for _, v := range arr {
		if _, ok := m[v]; ok {
			continue
		}
		res = append(res, v)
		m[v] = struct{}{}
	}

	fmt.Printf("subset: %v\n", res)
}
