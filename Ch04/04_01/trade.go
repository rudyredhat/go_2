// struct demo
// In go everything starts with starts with Capital letter is accessible outside
package main

import (
	"fmt"
)

// struct is used to define several fields
// Trade is a trade in stocks
type Trade struct {
	Symbol string  // Stock symbol
	Volume int     // Number of shares
	Price  float64 // Trade price
	Buy    bool    // true if buy trade, false if sell trade
}

func main() {
	// creating a Trade object & assign the values by position
	t1 := Trade{"MSFT", 10, 99.98, true}
	fmt.Println(t1)

	// %v for detailed and better view
	fmt.Printf("%+v\n", t1)

	// to access the individual fields we can use dots
	fmt.Println(t1.Symbol)

	// creating another object
	t2 := Trade{
		Symbol: "MSFT",
		Volume: 10,
		Price:  99.98,
		Buy:    true,
	}
	fmt.Printf("%+v\n", t2)

	// assigning 0 values to all the fields
	t3 := Trade{}
	fmt.Printf("%+v\n", t3)
}
