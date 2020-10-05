// channels = preferred way to communicate between go routines is by channels
// are like one directional pipe, send value from end and receive at other end
// two kind of channel = buffered and unbuffered
// buffered channels are some like bonded queues
package main

import (
	"fmt"
	"time"
)

func main() {
	// create a channel using make built in function
	ch := make(chan int)

	// This will block
	/*
		<-ch
		fmt.Println("Here")
	*/

	go func() {
		// Send number of the channel
		ch <- 353
	}()

	// Receive from the channel
	val := <-ch
	fmt.Printf("got %d\n", val)

	fmt.Println("-----")

	// Send multiple channels
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < 3; i++ {
		val := <-ch
		fmt.Printf("received %d\n", val)
	}

	fmt.Println("-----")

	// close to signal we're done
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	for i := range ch {
		fmt.Printf("received %d\n", i)
	}
}
