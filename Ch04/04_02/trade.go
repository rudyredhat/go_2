// Method demo on struct
// define a value method which will return how much money its worth

package main

import (
	"fmt"
)

// Trade is a trade in stocks
type Trade struct {
	Symbol string  // Stock symbol
	Volume int     // Number of shares
	Price  float64 // Trade price
	Buy    bool    // true if buy trade, false if sell trade
}

// Value returns the trade value
// we have a receiver which is pointer to Trade object
// but why we are using pointer receiver, we can look at point.go file
// then the method name and the return type
func (t *Trade) Value() float64 {
	value := float64(t.Volume) * t.Price
	// if its a buy rate we need to reverse the value
	if t.Buy {
		value = -value
	}

	return value
}

func main() {
	t := Trade{
		Symbol: "MSFT",
		Volume: 10,
		Price:  99.98,
		Buy:    true,
	}
	fmt.Println(t.Value())
}
