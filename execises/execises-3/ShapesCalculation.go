package main

import (
	"fmt"
	"math"
)

//Shape interface for Circle and Reactangel
type Shape interface {
	area() float64
}

//Circle struct
type Circle struct {
	x, y, r float64
}

//Rectangel strict
type Rectangel struct {
	x1, y1, x2, y2 float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (r *Rectangel) area() float64 {
	l := distance(r.x1, r.y1, r.x2, r.y1)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func main() {
	c := Circle{0, 0, 5}
	r := Rectangel{0, 0, 10, 10}

	fmt.Println("Area of the rectangle is: ", r.area())
	fmt.Println("Area of the Circle is: ", c.area())
	fmt.Println("Sum of areas : ", totalArea(&c, &r))
}
