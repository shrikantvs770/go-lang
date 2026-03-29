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
