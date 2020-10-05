// go has garbage collector, which means we do not have to deal with memory management
package main

import (
	"fmt"
)

// cleanup is a function which will free up the resource
func cleanup(name string) {
	fmt.Printf("Cleaning up %s\n", name)
}

// worker is the actual code
func worker() {
	// use a defer code to free this resource
	// called in the reverse order first B then A
	defer cleanup("A")
	defer cleanup("B")

	fmt.Println("worker")
}

func main() {
	worker()
}
