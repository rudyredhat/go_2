// Print a friendly greeting
/*
multi line comment and above is a single line comment
*/

// main package used for code compile and execute
package main

// importing fmt package for format and print
import (
	"fmt"
)

// main func with keyword: func & body is enclosed in curly braces
func main() {
	// new line is print with the stings encoded
	fmt.Println("Welcome Gophers ☺")
}

/*
 go run welcome.go
 now create a executable that can be distribute
 go build welcome.go
 a file called welcome is created and can be run using ./welcome
*/
