// go mod tool to handle the dependency management
// install one dependency at a time and install the latest ver of the package
/*
	1) we need to work with toml so we need to build one module
	2) go mod init cfg, where cfg is the name of the package
	3) currently it will have a name of the module which is newly created
	4) edit the file and add the requirement
	- module cfg
	- require github.com/pelletier/go-toml v1.2.0
	5) run the prog , go run cfg.go - first it will find the dependency and then it will run thr prog
	6) a new file is created go.sum = where the list of packages we actually require
	7) go.mod is the file we list the packages we want to install
*/
package main

import (
	"fmt"
	"log"
	"os"

	toml "github.com/pelletier/go-toml"
)

// Config is configuration
type Config struct {
	Login struct {
		User     string
		Password string
	}
}

func main() {
	file, err := os.Open("config.toml")
	if err != nil {
		log.Fatalf("error: can't open config file - %s", err)
	}
	defer file.Close()

	cfg := &Config{}
	dec := toml.NewDecoder(file)
	if err := dec.Decode(cfg); err != nil {
		log.Fatalf("error: can't decode configuration file - %s", err)
	}

	fmt.Printf("%+v\n", cfg)
}
