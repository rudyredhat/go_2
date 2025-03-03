/*
An "even ended number" is a number who's first and last digit are the same.
You mission, should you choose to accept it, is to count how many "even ended
numbers" are there that are a multiplication of 4 digit numebrs.
*/

package main

import (
	"fmt"
)

// can be easily checked if it converted to a string -  Sprinf
func main() {
	// count = 0
	count := 0

	// for every pair of 4 digit numbers
	for a := 1000; a <= 9999; a++ {
		for b := a; b <= 9999; b++ { // don't count twice
			n := a * b

			// if a * b even ended
			// convert int to string
			s := fmt.Sprintf("%d", n)
			// check the first digit at 0 is equal to the last digit
			if s[0] == s[len(s)-1] {
				// count = count + 1, increment the count
				count++
			}
		}
	}

	// print count
	fmt.Println(count)
}
