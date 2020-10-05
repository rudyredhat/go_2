package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `
	Needles and pins
	Needles and pins
	Sew me a sail
	To catch me the wind
	`
	// Fields func will remove the extra white space and will return only the string
	words := strings.Fields(text)
	// create a map and assign the string to the key and the value is the count of words
	counts := map[string]int{} // Empty map
	// we will iterate over the words , as we are not interested in the index
	for _, word := range words {
		println(word)
		/*
			we will increment the counter for the speicific word, first converting it
			to lower case and then ++
		*/
		counts[strings.ToLower(word)]++
	}

	fmt.Println(counts)
}
