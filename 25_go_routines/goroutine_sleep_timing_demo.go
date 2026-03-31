// Demonstrates goroutines with different execution times and timing their completion.
// Shows how main can use Sleep to ensure all goroutines finish, and measures total elapsed time.
package main

import (
	"fmt"
	"time"
)

func greet() {
	// this will wait for 5 seconds
	time.Sleep(time.Second * 5)
	fmt.Println("Hi from Go routine 2")
}

func main() {
	start := time.Now()
	// go routine 1
	go greet()

	// go routine 2
	go func() {
		// this will invoke immediately, and finish before go routine 1
		fmt.Println("Hi from Go routine 1")
	}()

	// wait some time so that the main wait for the go route to finish
	time.Sleep(time.Second * 10)

	// as main doesn't wait for Go routine to finish
	fmt.Println("Done!")

	fmt.Println("total time took ", time.Since(start)) // 10.004222665s Why? when main func goes is put to sleep in parallel go routine 1 also sleeps for 5 secs

}
