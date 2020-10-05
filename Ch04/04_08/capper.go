// implement io.Writer interface
// transforms everything to uppercase
package main

import (
	"fmt"
	"io"
	"os"
)

// Capper implements io.Writer and turns everything to uppercase
type Capper struct {
	wtr io.Writer
}

func (c *Capper) Write(p []byte) (n int, err error) {
	diff := byte('a' - 'A')

	// creating a alternative buffer using built in make function
	out := make([]byte, len(p))
	// iterate over the slice
	for i, c := range p {
		if c >= 'a' && c <= 'z' {
			c -= diff
		}
		out[i] = c
	}

	return c.wtr.Write(out)
}

func main() {
	c := &Capper{os.Stdout}
	fmt.Fprintln(c, "Hello there")
}
