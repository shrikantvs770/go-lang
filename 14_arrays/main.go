package main

import "fmt"

func main() {
	// array, fixed cannot grow
	var marks [3]int
	marks[0] = 78
	marks[2] = 45
	fmt.Println(marks)
	fmt.Println(len(marks))

	// arrays inline declaration
	weights := [5]int{56, 67, 89, 87, 100}
	fmt.Println(weights, len(weights))

	weights2 := [...]int{56, 67, 89, 87, 100} // this is an array brother. This is an array, like [5]int{…}, but Go automatically infers the length.
	fmt.Println(weights2) 
}
