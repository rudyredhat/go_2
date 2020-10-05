// Get content type of sites
// To handle the concurrency
// concurrency = is the composition of idependently executed processes
// parallelism = simultaneous execution
// query bunch of urls and check the ret type
package main

import (
	"fmt"
	"net/http"
	"sync" // use of WaitGroup
)

// func which does a http call and check the header and prints it out
func returnType(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	fmt.Printf("%s -> %s\n", url, ctype)
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/xml",
	}

	// create a WaitGroup
	var wg sync.WaitGroup
	for _, url := range urls {
		// for every check we need to tell we added a WaitGroup
		wg.Add(1)
		// using the go keyword to generate the go routines & create a anonympus func
		go func(url string) {
			returnType(url)
			// here we need to signal that we are done
			wg.Done()
		}(url)
	}
	// wait for all the go routine to finish
	wg.Wait()
}
