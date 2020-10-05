// JSON example
/*
	1) in distributed network we need to send and receive networks
	2) json = preferred encoding in rest APIs

*/
package main

import (
	"bytes"
	"encoding/json" // built in encoding .json
	"fmt"
	"log"
	"os"
)

var data = `
{
  "user": "Scrooge McDuck",
  "type": "deposit",
  "amount": 1000000.3
}
`

// Request is a bank transactions
type Request struct {
	// must start with upper case
	Login  string  `json:"user"`
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}

func main() {
	// create a buffer string
	rdr := bytes.NewBufferString(data) // Simulate a file/socket

	// Decode request
	dec := json.NewDecoder(rdr)

	// create an empty req
	req := &Request{}
	if err := dec.Decode(req); err != nil {
		log.Fatalf("error: can't decode - %s", err)
	}

	fmt.Printf("got: %+v\n", req)

	// Create response
	prevBalance := 8500000.0 // Loaded from database
	// here the values are the empty interfaces - in go means any type ***
	resp := map[string]interface{}{
		"ok":      true,
		"balance": prevBalance + req.Amount,
	}

	// Encode respose
	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(resp); err != nil {
		log.Fatalf("error: can't encode - %s", err)
	}
}
