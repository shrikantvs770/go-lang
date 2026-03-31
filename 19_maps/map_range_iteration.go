// Demonstrates iterating over a map using range in Go.
// Shows how to access both keys and values in a loop.
// Useful for processing all map entries.
package main

import "fmt"

func main() {
	stdMap := map[string]string{
		"u001": "Srikant V S",
		"u002": "Elias Samson",
		"u003": "WWE",
		"u004": "John Cena",
		"u005": "Shawn Michaels",
	}

	for key, val := range stdMap {
		fmt.Println(key, val)
	}
}
