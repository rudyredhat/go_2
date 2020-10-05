package main

import (
	"fmt"
)

// takes the slice of int and doubles the value at that index
func doubleAt(values []int, i int) {
	values[i] *= 2
}

func double(n int) {
	n *= 2
}

// this function gets a pointer integer
func doublePtr(n *int) {
	// here we are de-refrencing the pointer and assigning the value
	*n *= 2
}

func main() {
	// create a slice of value 1,2,3,4
	values := []int{1, 2, 3, 4}
	// double the value at index number 2
	// but passing the slice is done by refernce and changes would be affected
	doubleAt(values, 2)
	fmt.Println(values)

	val := 10
	/* 	passing an int to func, it passes by value
	means create a copy of the int and pass it to func
	any changes inside the func, won't affect the original value
	*/
	double(val)
	fmt.Println(val)
	// passing the pointer through a value
	doublePtr(&val)
	fmt.Println(val)
}
