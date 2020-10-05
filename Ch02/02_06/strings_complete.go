// Go strings
package main

import (
	"fmt"
)

func main() {
	book := "The colour of magic"
	fmt.Println(book)

	fmt.Println(len(book))

	fmt.Printf("book[0] = %v (type %T)\n", book[0], book[0])
	// book[0] = 84 (unit8) - o/p of last line
	// uint8 is a byte

	// Strings in go are immutable
	// book[0] = 116

	// Slice (start, end), 0 based, ½ empty range
	fmt.Println(book[4:11])

	// Slice (no end)
	fmt.Println(book[4:])

	// Slice (no start)
	fmt.Println(book[:4])

	// Use + to concatenate strings
	fmt.Println("t" + book[1:])

	// Unicode, we can insert 1/2 char inside the string
	fmt.Println("It was ½ price!")

	// Multi line string input
	poem := `
	The road goes ever on
	Down from the door where it began
	...
	`
	fmt.Println(poem)
}
