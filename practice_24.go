package main

import (
	"fmt"
	"math"
)

type point struct {
	x int
	y int
}

func newPoint(x, y int) *point {
	return &point{
		x: x,
		y: y,
	}
}

func (p *point) calcDistance(to *point) float64 {
	xD, yD := p.x-to.x, p.y-to.y
	return math.Sqrt(float64(xD*xD + yD*yD))
}

func main() {
	p1 := newPoint(10, 1)
	p2 := newPoint(5, 5)
	dist := p1.calcDistance(p2)
	fmt.Printf("Distance between p1 and p2: %0.4f", dist)
}
