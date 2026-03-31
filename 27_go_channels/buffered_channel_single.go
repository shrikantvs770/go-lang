// Shows how a buffered channel (size 1) prevents deadlock.
// Demonstrates sending and receiving a single value using a buffer.
package main

import (
	"fmt"
	"strings"
)

func main() {
	ping := make(chan string, 1) // 1 buffer man

	ping <- "hello" // this is blocking but can hold 1 value

	word := strings.ToUpper(<-ping) // can read from the buffer that value

	fmt.Printf("word: %v\n", word)
	word = strings.ToUpper(<-ping)

	fmt.Printf("word: %v\n", word)

}
