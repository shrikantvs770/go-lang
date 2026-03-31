// Demonstrates using time.Ticker to trigger periodic events in Go.
// Shows how to execute code at regular intervals using a blocking for loop.
package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.NewTicker(time.Second * 2) // every 2 sec

	for {
		current := <-tick.C // this will block the for loop
		fmt.Printf("current: %v\n", current)
	}
}
