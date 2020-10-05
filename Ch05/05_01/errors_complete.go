// pkg/errors example
// a prog reads a config and starts to execute as well saves in log file **imp**
/*
 	1) in py and java it will show us a exception why the file is not able to open,
	if the path is incorrect.
	2) so we need to use in terminal - go get github.com/pkg/errors
	3) when we want to see the full call stack , we can use +v format
	its better not show the stack tarce, otherwise users think its a bug in prog
	4) that's why we save stack traces in log file
*/
package main

import (
	"fmt"
	"log" // built in package for stack trace
	"os"

	// so we have imported it here
	"github.com/pkg/errors"
)

// Config holds configuration values
type Config struct {
	// configuration fields go here (redacted)
}

// which read the configuration from a file
// returns a pointer to the configuration object and a possible error
func readConfig(path string) (*Config, error) {
	// will try to open the configuration file
	file, err := os.Open(path)
	// else will print the error
	if err != nil {
		// so here we wrap the error and returning one msg
		return nil, errors.Wrap(err, "can't open configuration file")
	}

	// close the file once the work is done
	defer file.Close()

	cfg := &Config{}
	// Parse file here (redacted)
	return cfg, nil

}

// function for setting up the logging
func setupLogging() {
	// open a log file in append mode
	out, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	// and send output to this file
	log.SetOutput(out)
}

func main() {
	// call the logging
	setupLogging()
	cfg, err := readConfig("/path/to/config.toml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		// print the logs in the log file
		log.Printf("error: %+v", err)
		os.Exit(1)
	}

	// Normal operation (redacted)
	fmt.Println(cfg)
}
