package main

import "fmt"

type triangle struct {
	heigth float64
	base   float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	return t.heigth * t.base * 0.5
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func main() {
	t := triangle{5, 2}
	s := square{5}

	printArea(t)
	printArea(s)
}
