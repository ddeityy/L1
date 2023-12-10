package main

/*
Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point
с инкапсулированными параметрами x,y и конструктором.
*/

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

type Line float64

func NewPoint(x float64, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

func (p *Point) DistanceTo(p2 Point) float64 {
	return math.Sqrt((p2.x-p.x)*(p2.x-p.x) + (p2.y-p.y)*(p2.y-p.y))
}

func main() {
	p1 := NewPoint(1.1, 10.2)
	p2 := NewPoint(-9.2, 42.9)

	fmt.Printf("distance between points (%v, %v) and (%v, %v) is %.2f\n", p1.x, p1.y, p2.x, p2.y, p1.DistanceTo(p2))
}
