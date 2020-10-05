// Get content type of sites (using channels)
// instead of WaitGroup we are using here channels
package main

import (
	"fmt"
	"net/http"
)

// here is a extra out channel
func returnType(url string, out chan string) {
	resp, err := http.Get(url)
	if err != nil {
		// if err then send err to the channel
		out <- fmt.Sprintf("%s -> error: %s", url, err)
		return
	}

	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	// otherwise send the content type to the channel
	out <- fmt.Sprintf("%s -> %s", url, ctype)
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/xml",
	}

	// Create response channel
	ch := make(chan string)
	for _, url := range urls {
		go returnType(url, ch)
	}

	for range urls { // Run number of URLs times
		// get the result and print
		out := <-ch
		fmt.Println(out)
	}
}
