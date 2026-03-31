// Shows a blocking for loop using a time.Ticker channel.
// Continuously prints the current time every second in an infinite loop.
package main

import (
	"fmt"
	"time"
)

func main() {

	// blocking for loop
	ticker := time.NewTicker(time.Second)

	for {
		tick := <-ticker.C
		fmt.Println(tick)
	}

}
