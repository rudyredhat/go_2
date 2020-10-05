// Calculate md5 sum concurrently
package main

import (
	"bufio"
	"crypto/md5" // use of standard lib to cal md5
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// Parse the signature file, return a map of path->signature
func parseSignaturesFile(path string) (map[string]string, error) {
	// first we open the file and check for the error
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// create a map
	sigs := make(map[string]string)
	// use a scanner to go through the file
	scanner := bufio.NewScanner(file)
	for lnum := 1; scanner.Scan(); lnum++ {
		// ae5252a205000e972b9747b0b125cf96  nasa-05.log
		// split the lines into several fields and get the line from scanner.Text
		fields := strings.Fields(scanner.Text())
		if len(fields) != 2 {
			return nil, fmt.Errorf("%s:%d bad line", path, lnum)
		}
		// otherwise set for the sig for this file in the map
		sigs[fields[1]] = fields[0]
	}
	// end of the loop check for any error
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	// otherwise ret sig and nil val
	return sigs, nil
}

// func to get path(str) and ret a str
func fileMD5(path string) (string, error) {
	// open the file and check for the error
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

// result type from worker
type result struct {
	path  string
	match bool
	err   error
}

// run the go routine and process the file
func md5Worker(path string, sig string, out chan *result) {
	// create a result obj
	r := &result{path: path}
	s, err := fileMD5(path)
	if err != nil {
		r.err = err
		out <- r
		return
	}

	r.match = (s == sig)
	out <- r
}

func main() {
	sigs, err := parseSignaturesFile("md5sum.txt")
	if err != nil {
		log.Fatalf("error: can't read signaure file - %s", err)
	}

	out := make(chan *result)
	for path, sig := range sigs {
		go md5Worker(path, sig, out)
	}

	ok := true
	for range sigs {
		r := <-out
		switch {
		case r.err != nil:
			fmt.Printf("%s: error - %s\n", r.path, r.err)
			ok = false
		case !r.match:
			fmt.Printf("%s: signature mismatch\n", r.path)
			ok = false
		}
	}

	if !ok {
		os.Exit(1)
	}
}
