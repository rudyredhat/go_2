// select example
// build in select function allows to work with multiple channels together
package main

import (
	"fmt"
	"time"
)

func main() {
	// create 2 channels
	ch1, ch2 := make(chan int), make(chan int)

	// have a func that send the value to ch1
	go func() {
		ch1 <- 42
	}()

	select {
	case val := <-ch1:
		fmt.Printf("got %d from ch1\n", val)
	case val := <-ch2:
		fmt.Printf("got %d from ch2\n", val)
	}

	fmt.Println("----")
	// create timeouts
	out := make(chan float64)
	go func() {
		time.Sleep(100 * time.Millisecond)
		out <- 3.14
	}()

	select {
	case val := <-out:
		fmt.Printf("got %f\n", val)
	case <-time.After(20 * time.Millisecond):
		fmt.Println("timeout")
	}
}
