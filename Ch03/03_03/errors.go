// use for signal errors
// go has something like panic which is similar to exceptions, but not used in go
package main

import (
	"fmt"
	"math"
)

// calculate the square root of a number
// but will return error in the negative number
func sqrt(n float64) (float64, error) {
	if n < 0 {
		// return some value and fmt.Errorf will create a new error
		return 0.0, fmt.Errorf("sqrt of negative value (%f)", n)
	}
	// nil is just like null
	return math.Sqrt(n), nil
}

func main() {
	// positive number check
	s1, err := sqrt(2.0)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	} else {
		fmt.Println(s1)
	}

	// negative number check
	s2, err := sqrt(-2.0)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	} else {
		fmt.Println(s2)
	}
}
