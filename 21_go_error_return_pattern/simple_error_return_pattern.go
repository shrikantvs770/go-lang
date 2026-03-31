// Demonstrates a basic error return pattern in Go for input validation.
// Shows how to return and handle errors from a function.
package main

import (
	"fmt"
	"log"
)

func main() {
	_, err := isEligibleToVote(16)
	if err != nil {
		fmt.Println("Oops, ", err)
		log.Fatal(err)
	}
}

func isEligibleToVote(age int) (bool, error) {
	if age < 18 {
		return false, fmt.Errorf("Person is not eligible to vote.")
	} else {
		return true, nil // nil is fine means no error
	}
}
