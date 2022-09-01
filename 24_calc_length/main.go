package main

import (
	"fmt"
	"math"
)

type point struct {
	x float64
	y float64
}

func main() {
	first := newPoint(1, 1)
	second := newPoint(2, 2)

	fmt.Println(length(*first, *second))
}

func newPoint(x float64, y float64) *point {
	p := new(point)
	p.x = x
	p.y = y
	return p
}

func (p point) X() float64 {
	return p.x
}

func (p point) Y() float64 {
	return p.y
}

func length(first point, second point) float64 {
	return math.Sqrt((second.X()-first.X())*(second.X()-first.X()) + (second.Y()-first.Y())*(second.Y()-first.Y()))
}
