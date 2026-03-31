// Demonstrates basic map creation and usage in Go.
// Shows how to declare, initialize, and access map elements.
// Prints map contents and length.
package main

import (
	"fmt"
)

func main() {
	// Map : Key value pair
	ages := map[string]int{
		"srikant": 29,
		"John":    35,
	}
	fmt.Println(ages)
	fmt.Println("John Age", ages["John"])
	fmt.Println("len ", len(ages))
}
