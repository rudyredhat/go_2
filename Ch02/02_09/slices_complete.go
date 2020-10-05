// Go slices
package main

import (
	"fmt"
)

func main() {
	// Same type
	// definig the slice of strings with []string{"list of items"}
	loons := []string{"bugs", "daffy", "taz"}
	// %v will print out the list of slices
	fmt.Printf("loons = %v (type %T)\n", loons, loons)

	// Length
	fmt.Println(len(loons)) // 3

	fmt.Println("----")
	// 0 indexing
	fmt.Println(loons[1]) // daffy

	fmt.Println("----")
	// slices
	fmt.Println(loons[1:]) // [daffy taz]

	fmt.Println("----")

	// for loop to iterate through the slice of items
	for i := 0; i < len(loons); i++ {
		fmt.Println(loons[i])
	}

	fmt.Println("----")
	// Single value range , i = 0,1,2
	for i := range loons {
		fmt.Println(i)
	}

	fmt.Println("----")
	// Double value range, with both i and the name
	for i, name := range loons {
		fmt.Printf("%s at %d\n", name, i)
	}

	fmt.Println("----")
	// Double value range, ignore index by using _
	// because unused values are compilation error
	for _, name := range loons {
		fmt.Println(name)
	}

	fmt.Println("----")
	// append the slice of items
	loons = append(loons, "elmer")
	fmt.Println(loons) // [bugs daffy taz elmer]
}
