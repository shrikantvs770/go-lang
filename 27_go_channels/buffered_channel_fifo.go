// Demonstrates FIFO behavior with a buffered channel (size 3).
// Multiple values are sent and received in order, showing how buffering works.
package main

import (
	"fmt"
	"strings"
)

func main() {
	ping := make(chan string, 3) // 3 buffer man

	ping <- "amar" // this blocking but can hold 1 value
	ping <- "akbar"
	ping <- "anthony"

	// They will come in FIFO order

	word := strings.ToUpper(<-ping) // can read from the buffer

	fmt.Printf("word: %v\n", word)
	word = strings.ToUpper(<-ping)

	fmt.Printf("word: %v\n", word)

	word = strings.ToUpper(<-ping)
	fmt.Printf("word: %v\n", word)

}
