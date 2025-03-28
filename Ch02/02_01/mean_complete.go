// Calculate the mean of two numbers
package main

import (
	"fmt"
)

func main() {
	/*
		var x int
		var y int

		x = 1
		y = 2
	*/
	x, y := 1.0, 2.0

	// %v will print the go object and %T type
	fmt.Printf("x=%v, type of %T\n", x, x)
	fmt.Printf("y=%v, type of %T\n", y, y)

	// var mean int
	// mean = (x+y)/2
	mean := (x + y) / 2.0
	fmt.Printf("result: %v, type of %T\n", mean, mean)
}
