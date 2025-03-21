// Writing a function that return Content-Type header
package main

import (
	"fmt"
	"net/http"
)

// contentType will return the value of Content-Type header returned by making an
// HTTP GET request to url
func contentType(url string) (string, error) {
	// making an http get request, which will give response and error
	resp, err := http.Get(url)
	// if not nil, means some error has happen
	if err != nil {
		return "", err
	}

	// make sure to close the response body
	defer resp.Body.Close() // Make sure we close the body

	// get the contentType Header
	ctype := resp.Header.Get("Content-Type")
	if ctype == "" { // Return error if Content-Type header not found
		return "", fmt.Errorf("can't find Content-Type header")
	}

	// return Content type and nil , if there is no error
	return ctype, nil
}

func main() {
	ctype, err := contentType("https://linkedin.com")
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	} else {
		fmt.Println(ctype)
	}
}
