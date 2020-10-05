// function to killserver that reads a process identifier from pidFile
// converts to an int and print killing pid
package main

import (
	"fmt"
	"io/ioutil" // to read a file
	"log"
	"os"
	"strconv" // to convert the file content to int
	"strings"

	"github.com/pkg/errors" // to wrap error
)

// get a pidFile as a string and returns the error
func killServer(pidFile string) error {
	// tries to read the file
	data, err := ioutil.ReadFile(pidFile)
	if err != nil {
		// if err wrap the err
		return errors.Wrap(err, "can't open pid file (is server running?)")
	}

	if err := os.Remove(pidFile); err != nil {
		// We can go on if we fail here
		log.Printf("warning: can't remove pid file - %s", err)
	}

	// string(data) to wrap the data and strings.TrimSpace to trim all the spaces
	strPID := strings.TrimSpace(string(data))
	// convert the pid from a str to int
	pid, err := strconv.Atoi(strPID)
	if err != nil {
		return errors.Wrap(err, "bad process ID")
	}

	// Simulate kill
	fmt.Printf("killing server with pid=%d\n", pid)
	return nil
}

func main() {
	if err := killServer("server.pid"); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
