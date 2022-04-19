// Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point с инкапсулированными
// параметрами x,y и конструктором.
//
package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p Point) Distance(otherPoint Point) float64 {
	return math.Sqrt(math.Pow(p.x-otherPoint.x, 2) + math.Pow(p.y-otherPoint.y, 2))
}

func NewPoint(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

func main() {
	center := NewPoint(0,0)

	p := NewPoint(2,1)

	fmt.Printf("distance: %.2f\n", p.Distance(center))
}
