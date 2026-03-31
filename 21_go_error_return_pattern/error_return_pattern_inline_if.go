// Shows error handling using an inline if statement for concise error checks.
// Demonstrates returning and handling errors in a single line.
package main

import (
	"fmt"
	"log"
)

func main() {

	if _, err := isEligibleToVote(18); err != nil {
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
