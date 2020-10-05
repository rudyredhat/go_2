/*
	1) Usual method in go to signaling errors is to return the error value
	2) there are cases when we can not handle the error and try to signal it
	3) in go we have built-in panic function
	4) its helpful to rettun the stack traces
	5) Use of panic is discouraged in go *****not use panic*****
*/
package main

import (
	"fmt"
	//"os"
)

// create a func which takes slice of values and index and ret int
func safeValue(vals []int, index int) int {
	// if we defer a func
	defer func() {
		// recover is used to wrap untrusted code, such as third-party libraries
		if err := recover(); err != nil {
			fmt.Printf("ERROR: %s\n", err)
		}
	}()

	return vals[index]
}

func main() {
	/*
		vals := []int{1, 2, 3}
		// This will cause a panic and give index out of range
		v := vals[10]
		fmt.Println(v)
	*/
	/*
		file, err := os.Open("no-such-file")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		fmt.Println("file opened")
	*/

	v := safeValue([]int{1, 2, 3}, 10)
	fmt.Println(v)
}
