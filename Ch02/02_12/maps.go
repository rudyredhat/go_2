// Go's map data structure, where keys points to specific value
// keys must be of same type and values must be of same type
package main

import (
	"fmt"
)

func main() {
	// map strings to float
	stocks := map[string]float64{
		"AMZN": 1699.8,
		"GOOG": 1129.19,
		"MSFT": 98.61, // Must have trailing comma in multi line
	}

	// Number of items - it will print 3 = only count keys
	fmt.Println(len(stocks))

	// Get a value
	fmt.Println(stocks["MSFT"])

	// Get zero value if not found
	fmt.Println(stocks["TSLA"]) // 0

	// Use two value to see if found, value will get  the associated values
	// and ok will get true and false as per the value returned
	value, ok := stocks["TSLA"]
	if !ok {
		fmt.Println("TSLA not found")
	} else {
		fmt.Println(value)
	}

	// Set a new value
	stocks["TSLA"] = 322.12
	fmt.Println(stocks)

	// Delete a value, use delete func
	delete(stocks, "AMZN")
	fmt.Println(stocks)

	// Single value "for" is on keys
	for key := range stocks {
		fmt.Println(key)
	}

	// Double value "for" is key, value
	for key, value := range stocks {
		fmt.Printf("%s -> %.2f\n", key, value)
	}
}
