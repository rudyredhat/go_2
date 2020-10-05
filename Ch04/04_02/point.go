// Receiver example
package main

import (
	"fmt"
)

// Point is a 2d point
type Point struct {
	X int
	Y int
}

// Move moves the point with delta x and delta y
// if we are not using pointer reciever , will get a copy of the Move method
// so we use a pointer receiver
func (p *Point) Move(dx int, dy int) {
	p.X += dx
	p.Y += dy
}

func main() {
	// here we create a point and move it
	// creating a pointer to a point
	p := &Point{1, 2}
	p.Move(2, 3)
	fmt.Printf("%+v\n", p)
}
