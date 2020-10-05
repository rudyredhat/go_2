// "for" loop examples
package main

import (
	"fmt"
)

func main() {
	// basic for loop from 0,1,2
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	fmt.Println("----")
	for i := 0; i < 3; i++ {
		// use of break in if statement
		if i > 1 {
			break
		}
		fmt.Println(i)
	}

	fmt.Println("----")
	for i := 0; i < 3; i++ {
		// use of continue to skip the statement
		if i < 1 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("----")
	// just similar to while loop using a certain condition iterate through loop
	a := 0
	for a < 3 {
		fmt.Println(a)
		a++
	}

	fmt.Println("----")
	// similar to while True in other languages
	b := 0
	for {
		if b > 2 {
			break
		}
		fmt.Println(b)
		b++
	}
}
