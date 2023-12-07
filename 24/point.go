package main

import "math"

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
