// Method demo
// Use of Constructor / Initializer when an object is created
package main

import (
	"fmt"
	"os"
)

// Trade is a trade in stocks
type Trade struct {
	Symbol string  // Stock symbol
	Volume int     // Number of shares
	Price  float64 // Trade price
	Buy    bool    // true if buy trade, false if sell trade
}

// NewTrade will create a new trade and will validate the input
// we create a function which is starting with New, returns a pointer to the created object
func NewTrade(symbol string, volume int, price float64, buy bool) (*Trade, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol can't be empty")
	}

	if volume <= 0 {
		return nil, fmt.Errorf("volume must be >= 0 (was %d)", volume)
	}

	if price <= 0.0 {
		return nil, fmt.Errorf("price must be >= 0 (was %f)", price)
	}

	// creating obj inside the function is easily handeled by go
	// and know to allocate trade on the heap
	// create a pointer to the Trade object
	trade := &Trade{
		Symbol: symbol,
		Volume: volume,
		Price:  price,
		Buy:    buy,
	}
	// return the trade object and nil signifying no error
	return trade, nil
}

// Value returns the trade value
func (t *Trade) Value() float64 {
	value := float64(t.Volume) * t.Price
	if t.Buy {
		value = -value
	}

	return value
}

func main() {
	// create a Trade with a NewTrade
	t, err := NewTrade("MSFT", 10, 99.98, true)

	if err != nil {
		fmt.Printf("error: can't create trade - %s\n", err)
		os.Exit(1)
	}

	fmt.Println(t.Value())
}
