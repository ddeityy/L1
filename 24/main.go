package main

import "fmt"

/*
Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point
с инкапсулированными параметрами x,y и конструктором.
*/

func main() {
	p1 := NewPoint(1.1, 10.2)
	p2 := NewPoint(-9.2, 42.9)

	fmt.Printf("distance between points (%v, %v) and (%v, %v) is %.2f\n", p1.x, p1.y, p2.x, p2.y, p1.DistanceTo(p2))
}
