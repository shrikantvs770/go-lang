package main

import (
	"fmt"
)

func main() {

	// anonmymous func
	sumAll := func(nums ...int) int {

		sum := 0
		for _, val := range nums {
			sum += val
		}
		return sum
	}

	fmt.Printf("sumAll: %v\n", sumAll([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}...)) // its a slice



	// iife kind of
	func() {
		fmt.Println("Greetings")
	}()

}
