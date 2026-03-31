// Demonstrates a deadlock scenario with an unbuffered channel.
// The main goroutine tries to send without a concurrent receiver, causing a fatal deadlock.
package main

import (
	"fmt"
	"strings"
)

func main() {
	ping := make(chan string) // no buffer man

	ping <- "hello" // this is blocking and as there is no space in buffer, it expects the sender to be ready immediately, it will cause deadlock brother

	word := strings.ToUpper(<-ping)
	fmt.Printf("word: %v\n", word)

	// fatal error: all goroutines are asleep - deadlock!

}
