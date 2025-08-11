package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

func (p Point) Distance(other Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))
}

func main() {
	p1 := NewPoint(1.0, 2.0)
	p2 := NewPoint(4.0, 6.0)
	distance := p1.Distance(p2)

	fmt.Printf("Точка 1: (%.1f, %.1f)\n", p1.x, p1.y)
	fmt.Printf("Точка 2: (%.1f, %.1f)\n", p2.x, p2.y)
	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}