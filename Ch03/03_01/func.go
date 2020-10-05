// Basic function definition
package main

import (
	"fmt"
)

// add adds a to b
// func fun_name(var_name type) ret_type{}
func add(a int, b int) int {
	return a + b
}

// divmod returns quotient and reminder
// return multiple values
func divmod(a int, b int) (int, int) {
	return a / b, a % b
}

func main() {
	val := add(1, 2)
	fmt.Println(val)

	div, mod := divmod(7, 2)
	fmt.Printf("div=%d, mod=%d\n", div, mod)
}
