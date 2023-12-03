package main

import "fmt"

/*
Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов.
Последовательность в подмножноствах не важна.
Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/

func groupTempsByTen(temps []float32) map[int][]float32 {
	buckets := make(map[int][]float32)
	for _, v := range temps {
		k := int(v/10) * 10
		buckets[k] = append(buckets[k], v)
	}
	return buckets
}

func main() {
	temps := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	tempMap := groupTempsByTen(temps)
	for k, v := range tempMap {
		fmt.Println(k, v)
	}

}
