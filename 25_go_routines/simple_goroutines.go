// Demonstrates launching multiple goroutines in Go for concurrent execution.
// Shows how main can wait using Sleep to allow goroutines to finish.
package main

import (
	"fmt"
	"time"
)

func greet() {
	fmt.Println("Hi from Go routine 2")
}

func main() {

	// go routine 1
	go greet()

	// go routine 2
	go func() {
		fmt.Println("Hi from Go routine 1")
	}()

	time.Sleep(time.Second * 2)

	// as main doesn't wait for Go routine to finish
	fmt.Println("Done!")

}
