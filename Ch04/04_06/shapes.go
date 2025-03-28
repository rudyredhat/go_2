package main

import (
	"fmt"
	"math"
)

// Square is a square
type Square struct {
	Length float64
}

// Area returns the area of the square
func (s *Square) Area() float64 {
	return s.Length * s.Length
}

// Circle is a circle
type Circle struct {
	Radius float64
}

// Area returns the curcle of the square
func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// sumAreas return the sum of all areas in the slice
func sumAreas(shapes []Shape) float64 {
	total := 0.0

	for _, shape := range shapes {
		total += shape.Area()
	}

	return total
}

// Shape is a shape interface
// an interface is a collection of methods and each type implements all the methods
// in the interface that are satisfying the interface
type Shape interface {
	// method must match the name and the returning value
	Area() float64
}

func main() {
	s := &Square{20}
	fmt.Println(s.Area())

	c := &Circle{10}
	fmt.Println(c.Area())

	// slice of shape which have both square and circle and prints the total area
	shapes := []Shape{s, c}
	sa := sumAreas(shapes)
	fmt.Println(sa)
}
