// fmt.Sprint example

package main

import (
	"fmt"
)

func main() {
	n := 42
	// returns a formated string, by converting int to string
	s := fmt.Sprintf("%d", n)

	// %q = verb (q) which will distinguish the string and int and print as "42"
	// instead if we use %s, it will only print 42
	fmt.Printf("s = %q (type %T)\n", s, s)
}
