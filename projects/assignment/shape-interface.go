package main

import (
	"fmt"
)

type square struct {
	sideLength float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

type triangle struct {
	height float64
	base   float64
}

func (t triangle) getArea() float64 {
	return t.base * 0.5 * t.height
}

type shape interface {
	getArea() float64
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
