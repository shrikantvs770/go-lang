// Demonstrates deleting elements from a map in Go.
// Shows how to remove a key-value pair and print the result.
// Useful for understanding in-place map modifications.
package main

import "fmt"

func main() {
	var stdMap map[string]string = map[string]string{
		"u001": "Srikant V S",
		"u002": "Elias Samson",
		"u003": "WWE",
	}

	fmt.Println(stdMap)
	delete(stdMap, "u003") // it deletes in place brother
	fmt.Println(stdMap)
}
